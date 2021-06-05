package classfile

import "fmt"

// ClassFile store the parsed class file
type ClassFile struct {
	//magic uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

// Parse parses class file to ClassFile object
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}
func (cf *ClassFile) read(reader *ClassReader) {
	cf.readAndCheckMagic(reader)               // 见 3.2.3
	cf.readAndCheckVersion(reader)             // 见 3.2.4
	cf.constantPool = readConstantPool(reader) // 见 3.3
	cf.accessFlags = reader.readUint16()
	cf.thisClass = reader.readUint16()
	cf.superClass = reader.readUint16()
	cf.interfaces = reader.readUint16s()
	cf.fields = readMembers(reader, cf.constantPool) // 见 3.2.8
	cf.methods = readMembers(reader, cf.constantPool)
	cf.attributes = readAttributes(reader, cf.constantPool) //见 3.4
}

// MajorVersion exposes the major version of this class
func (cf *ClassFile) MajorVersion() uint16 {
	return cf.majorVersion
}

// ClassName get the class name from constant pool
func (cf *ClassFile) ClassName() string {
	return cf.constantPool.getClassName(cf.thisClass)
}

// SuperClassName get the super class name from constant pool
func (cf *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(cf.superClass)
	}
	return "" // does not have super class -> only has java.lang.Object
}

// InterfaceNames get the interface class name from constant pool
func (cf *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(cf.interfaces))
	for i, cpIndex := range cf.interfaces {
		interfaceNames[i] = cf.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

// readAndCheckMagic checks whether this class file starts with cafebabe
func (cf *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

//readAndCheckVersion check if the class file belongs to the supported version
func (cf *ClassFile) readAndCheckVersion(reader *ClassReader) {
	cf.minorVersion = reader.readUint16()
	cf.majorVersion = reader.readUint16()
	switch cf.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

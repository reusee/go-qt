package smoke_info

type Class struct {
  name string
  external bool
  parents []*Class
  size int
  has_constructor bool
  has_copy_constructor bool
  has_virtual_destructor bool
  is_namespace bool
  is_defined_elsewhere bool
}

const (
  T_VOIDP = 0
  T_BOOL
  T_CHAR
  T_UCHAR
  T_SHORT
  T_USHORT
  T_INT
  T_UINT
  T_LONG
  T_ULONG
  T_FLOAT
  T_DOUBLE
  T_ENUM
  T_CLASS

  KIND_STACK = iota
  KIND_POINTER
  KIND_REFERENCE
  KIND_NONE
)

type Type struct {
  name string
  class *Class
  typeId int
  kind int
  is_const bool
}

type Method struct {
  class *Class
  name string
  arguments []*Type
  ret *Type
  is_static bool
  is_const bool
  is_copy_constructor bool
  is_internal bool
  is_enum bool
  is_constructor bool
  is_destructor bool
  is_protected bool
  is_field_accessor bool
  is_property_accessor bool
  is_virtual bool
  is_purevirtual bool
  is_signal bool
  is_slot bool
  is_explicit bool
}

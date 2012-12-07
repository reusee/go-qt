package smoke_info

type Class struct {
  Name string
  External bool
  Parents []*Class
  Size int
  HasConstructor bool
  HasCopyConstructor bool
  HasVirtualDestructor bool
  IsNamespace bool
  IsDefinedElsewhere bool
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
  Name string
  Class *Class
  TypeId int
  Kind int
  IsConst bool
}

type Method struct {
  Class *Class
  Name string
  Arguments []*Type
  Ret *Type
  IsStatic bool
  IsConst bool
  IsCopyConstructor bool
  IsInternal bool
  IsEnum bool
  IsConstructor bool
  IsDestructor bool
  IsProtected bool
  IsFieldAccessor bool
  IsPropertyAccessor bool
  IsVirtual bool
  IsPurevirtual bool
  IsSignal bool
  IsSlot bool
  IsExplicit bool
}

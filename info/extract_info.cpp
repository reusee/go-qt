#include <smoke.h>
#include <smoke/qtcore_smoke.h>
#include <smoke/qtgui_smoke.h>

#include <iostream>
#include <fstream>
using namespace std;

void write_header(ofstream&);
void write_footer(ofstream&);
void generate_class_def(ofstream&, Smoke::Class);
void generate_class_inheritance(ofstream&, Smoke*, Smoke::Class);
void generate_type_info(ofstream&, Smoke*, int, Smoke::Type);
void generate_method_info(ofstream&, Smoke*, Smoke::Method, int);

int main(int argc, char *argv[]) {
  init_qtcore_Smoke();
  init_qtgui_Smoke();

  ofstream out;
  out.open("info.go");

  write_header(out);

  Smoke* smokes[] = {
    qtcore_Smoke,
    qtgui_Smoke,
  };

  for (auto i = begin(smokes); i != end(smokes); i++) {
    Smoke* smoke = *i;

    for (int i = 0; i < smoke->numClasses; ++i) {
      Smoke::Class klass = smoke->classes[i];
      if (!klass.className) continue;
      generate_class_def(out, klass);
    }

    for (int i = 0; i < smoke->numClasses; ++i) {
      Smoke::Class klass = smoke->classes[i];
      if (!klass.className) continue;
      generate_class_inheritance(out, smoke, klass);
    }

    for (int i = 0; i < smoke->numTypes; ++i) {
      Smoke::Type t = smoke->types[i];
      generate_type_info(out, smoke, i, t);
    }

    for (int i = 0; i < smoke->numMethods; ++i) {
      Smoke::Method method = smoke->methods[i];
      generate_method_info(out, smoke, method, i);
    }

  }

  write_footer(out);
  out.close();
  return 0;
}

void write_header(ofstream &out) {
  out << "package info" << endl;
  out << endl;
  out << "var classes = make(map[string]*Class)" << endl;
  out << "var types = make(map[int]*Type)" << endl;
  out << "var methods = make(map[int]*Method)" << endl;
  out << endl;
  out << "var klass *Class" << endl;
  out << "var type_ *Type" << endl;
  out << "var method *Method" << endl;
  out << endl;
  out << "func init() {" << endl;
}

void write_footer(ofstream &out) {
  out << "}" << endl;
}

void generate_class_def(ofstream& out, Smoke::Class klass) {
  string name(klass.className);
  out << "klass = &Class{" << endl;
  out << "\t\"" << name << "\"," << endl;
  if (klass.external)
    out << "\ttrue," << endl;
  else
    out << "\tfalse," << endl;
  out << "\tmake([]*Class, 0)," << endl;
  out << "\t" << klass.size << "," << endl;
  if (klass.flags & Smoke::cf_constructor)
    out << "\ttrue," << endl;
  else
    out << "\tfalse," << endl;
  if (klass.flags & Smoke::cf_deepcopy)
    out << "\ttrue," << endl;
  else
    out << "\tfalse," << endl;
  if (klass.flags & Smoke::cf_virtual)
    out << "\ttrue," << endl;
  else
    out << "\tfalse," << endl;
  if (klass.flags & Smoke::cf_namespace)
    out << "\ttrue," << endl;
  else
    out << "\tfalse," << endl;
  if (klass.flags & Smoke::cf_undefined)
    out << "\ttrue," << endl;
  else
    out << "\tfalse," << endl;
  out << "}" << endl;
  out << "classes[\"" << name << "\"] = klass" << endl;
}

void generate_class_inheritance(ofstream& out, Smoke* smoke, Smoke::Class klass) {
  Smoke::Index* idx = smoke->inheritanceList + klass.parents;
  string name(klass.className);
  while (*idx) {
    Smoke::Class parent = smoke->classes[*idx];
    if (!parent.className) continue;
    string parent_name(parent.className);
    out << "classes[\"" << name << "\"].parents = ";
    out << "append(classes[\"" << name << "\"].parents, classes[\"";
    out << parent_name << "\"])" << endl;
    idx++;
  }
}

void generate_type_info(ofstream& out, Smoke* smoke, int idx, Smoke::Type t) {
  if (!t.name) return;
  out << "type_ = &Type{" << endl;
  out << "\"" << t.name << "\"," << endl;
  if (t.classId != -1) {
    Smoke::Class klass = smoke->classes[t.classId];
    if (!klass.className) {
      out << "nil," << endl;
    } else {
      string name(klass.className);
      out << "classes[\"" << name << "\"]," << endl;
    }
  } else
    out << "nil," << endl;
  const char* typeIds[] = {
    "T_VOIDP", "T_BOOL", "T_CHAR", "T_UCHAR", "T_SHORT",
    "T_USHORT", "T_INT", "T_UINT", "T_LONG", "T_ULONG",
    "T_FLOAT", "T_DOUBLE", "T_ENUM", "T_CLASS",
  };
  out << typeIds[(t.flags & Smoke::tf_elem)] << "," << endl;
  if (t.flags & Smoke::tf_stack)
    out << "KIND_STACK," << endl;
  else if (t.flags & Smoke::tf_ptr)
    out << "KIND_POINTER," << endl;
  else if (t.flags & Smoke::tf_ref)
    out << "KIND_REFERENCE," << endl;
  else
    out << "KIND_NONE," << endl;
  if (t.flags & Smoke::tf_const)
    out << "true," << endl;
  else
    out << "false," << endl;
  out << "}" << endl;
  out << "types[" << idx << "] = type_" << endl;
}

void generate_method_info(ofstream& out, Smoke* smoke, Smoke::Method method, int idx) {
  Smoke::Class klass = smoke->classes[method.classId];
  if (!klass.className) return;
  string class_name(klass.className);
  out << "method = &Method{" << endl;
  out << "classes[\"" << class_name << "\"]," << endl;
  out << "\"" << smoke->methodNames[method.name] << "\"," << endl;
  out << "make([]*Type, 0)," << endl;
  out << "types[" << method.ret << "]," << endl;
  out << ((method.flags & Smoke::mf_static) ? "true" : "false") << "," << endl;
  out << ((method.flags & Smoke::mf_const) ? "true" : "false") << "," << endl;
  out << ((method.flags & Smoke::mf_copyctor) ? "true" : "false") << "," << endl;
  out << ((method.flags & Smoke::mf_internal) ? "true" : "false") << "," << endl;
  out << ((method.flags & Smoke::mf_enum) ? "true" : "false") << "," << endl;
  out << ((method.flags & Smoke::mf_ctor) ? "true" : "false") << "," << endl;
  out << ((method.flags & Smoke::mf_dtor) ? "true" : "false") << "," << endl;
  out << ((method.flags & Smoke::mf_protected) ? "true" : "false") << "," << endl;
  out << ((method.flags & Smoke::mf_attribute) ? "true" : "false") << "," << endl;
  out << ((method.flags & Smoke::mf_property) ? "true" : "false") << "," << endl;
  out << ((method.flags & Smoke::mf_virtual) ? "true" : "false") << "," << endl;
  out << ((method.flags & Smoke::mf_purevirtual) ? "true" : "false") << "," << endl;
  out << ((method.flags & Smoke::mf_signal) ? "true" : "false") << "," << endl;
  out << ((method.flags & Smoke::mf_slot) ? "true" : "false") << "," << endl;
  out << ((method.flags & Smoke::mf_explicit) ? "true" : "false") << "," << endl;
  out << "}" << endl;
  for (int i = 0; i < method.numArgs; i++) {
    out << "method.arguments = append(method.arguments, types[";
    out << smoke->argumentList[method.args + i] << "])" << endl;
  }
  out << "methods[" << idx << "] = method" << endl;
}

import excalibur
@excalibur.debug_with_prefix(">>>")
def print1(input1):
    print(input1)

print1(123)

@excalibur.debug_class
class Foo:
    def print1(self,input1):
        print(input1)
    def print2(self,input1):
        print(input1)

f = Foo()
f.print1("234")
f.print2("345")

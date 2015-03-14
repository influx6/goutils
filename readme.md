#GoUtils
GoUtils was created to store certain useful utility functions and pattern that are useful but redundantly rewritten constantly in all my golang packages. Its really simple.

##Install

      go get github.com/influx6/goutils 
Then

      go install github.com/influx6/goutils

##Utilites

 - Equal(a,b interface{}) bool
    Using abit of reflection it checks equality between golangs base types int,string,..etc


 - TypeCallers
    I needed a means of checking basic types but more than that to be notified by a callback depending on the basic type without having to write long lines of type assertions. TypeCallers is a struct with a attributes named after its basic type and each is a function which will be called when the data given to it matches or else the `unknown` caller is called.

    
            type TypeCallers struct {
                Int     func(int)
                UInt    func(uint)
                Int8    func(int8)
                UInt8   func(uint8)
                Int16   func(int16)
                UInt16  func(uint16)
                Int32   func(int32)
                UInt32  func(uint32)
                Int64   func(int64)
                UInt64  func(uint64)
                String  func(string)
                Byte    func(byte)
                Bytes   func([]byte)
                Float64 func(float64)
                Float32 func(float32)
                Unknown func(interface{})
            }
        
        
    


 - NewStringMorpher() *StringMorpher 
    StringMorpher creates a struct called the stringmorpher that standardizes what NewStringConverter does but allows a better synchronouze method call like below:

        
        
            morph := NewStringMorpher()

            val := morph.Morph(200) //returns the string version of the number 200

        
    

 - NewByteMorpher() *ByteMorpher 
    Heck whats the use of being able to turn anything possible stringable type into a string with StringMorpher if you cant turn that into a byte along the side. This creates a new Morpher off of StringMorpher and takes all the result of StringMorpher and transforms them into []bytes. You can say it composes `StringMorpher` into a `ByteMorpher`

        
            morph := NewByteMorpher()

            val := morph.Morph(200) //returns the  []byte version of the number 200



 - NewStringConverter(val *String) *TypeCallers 
    To make life easier this method auto-generates a TypeCaller for converting all the basic type into their string version,when it encounters the unknown state,it uses the json:encoder to marshall it incase if its a struct and if this fails well...(sorry ,no string for you :p) .
        

 - OnType(a interface}, caller *TypeCallers) 
    This method takes a value of interface{} and a TypeCaller and tells the caller to call the appropriate function based on the inferred type

 - IsBasicType(a interface}) bool 
        This method passes its value to all available ValueMorphers and returns true or false if it is a basic types

 - ValueMorphers
    ValueMorphers are cool,they are really functions of all available golang basic types and each will convert a value passed to it into its own type or else return two values(zeroth value of that type and an error)

        
            func ByteListMorph(a interface}) ([]byte, error) 

            func ByteMorph(a interface}) (byte, error) 

            func StringMorph(a interface}) (string, error) 

            func Float32Morph(a interface}) (float32, error) 

            func Float64Morph(a interface}) (float64, error) 

            func UInt16Morph(a interface}) (uint16, error) 

            func UInt32Morph(a interface}) (uint32, error) 

            func UInt64Morph(a interface}) (uint64, error) 

            func UIntMorph(a interface}) (uint, error) 

            func UInt8Morph(a interface}) (uint8, error) 

            func Int16Morph(a interface}) (int16, error) 
            func Int32Morph(a interface}) (int32, error) 

            func Int64Morph(a interface}) (int64, error) 

            func IntMorph(a interface}) (int, error) 

            func Int8Morph(a interface}) (int8, error) 


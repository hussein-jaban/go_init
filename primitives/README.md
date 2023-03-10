## SUMMARY

- ### Boolean type

  - Values are true or false
  - Not an alias for other types (e.g. int)
  - Zero value is false

- ### Numeric types

  - Integers
    - Signed integers
      - int type has varying size, but min 32 bits
      - 8 bit (int8) through 64 bit (int64)
    - Unsigned integers
      - 8 bit (byte and uint8) through 32 bit (uint32)
    - Arithmetic operations
      - Addition, subtraction, multiplication, division, remainder
    - Bitwise operations
      - And, or, xor, and not
    - Zero value is 0
    - Can't mix types in same family! (uint16 + uint32 = error)

- ### Floating point numbers

  - Follow IEEE-754 standard
  - Zero value is 0
  - 32 and 64 bit versions
  - Literal styles
    - Decimal (3.14)
    - Exponential (13e18 or 2E10)
    - Mixed (13.7e12)
  - Arithmetic operations
    - Addition, subtraction, multiplication, division

- ### Text types

  - Strings
    - UTF-8
    - Immutable
    - . Can be concatenated with plus (+) operator
    - Can be converted to []byte
  - Rune
    - UTF-32
    - Alias for int32
    - Special methods normally required to process
      - e.g. strings.Reader#ReadRune

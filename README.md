# Util Password

## Install

```
go get github.com/kuromittsu/util_password
```

## Password Hash

### Available Methods

```go
// convert string to hashed version
PasswordHash(rawPassword string) (string, error)

// compare hashed version using unhashed version
PasswordCompare(hashedPassword, rawPassword string) error
```

### Usage

#### Example

```go
// import utpass "github.com/kuromittsu/util_password"

tr := "password" // text right
tw := "passwordddd" // text wrong

trHashed, err := utpass.PasswordHash(tr)
if err != nil {
  log.Fatalf("err | %v", err)
}

fmt.Printf("trHashed: %v\n", trHashed)

err = util_password.PasswordCompare(trHashed, tr)
if err != nil {
  fmt.Printf("comparing should right | err: %v\n", err)
} else {
  fmt.Println("result comparing is right")
}

err = util_password.PasswordCompare(trHashed, tw)
if err != nil {
  fmt.Printf("comparing should wrong | err: %v\n", err)
} else {
  fmt.Println("result comparing is wrong")
}
```

Result

```txt
trHashed: $2a$10$lxGADOD9tsAA9N/8CEtD5.oVJ3kzZHI0ZiP5TyKm8YXNqFXKWo4C6 // different every run
result comparing is right
comparing should wrong | err: hash and password not match
```

## Password Validation

### Available Methods

```go
// validate text using validation rule(s) and language (ID/EN)
PasswordValidate(text string, rules []validateKey, lang Lang) passwordValidateResult

// default rules
DefaultRules() []validateKey
```

### Validation Rules

```go
// import utpass "github.com/kuromittsu/util_password"

// value minimum 8 length
utpass.MIN_8_LENGTH

// value minimum has 1 uppercase letter
utpass.MIN_1_UPPER

// value minimum has 1 lowercase letter
utpass.MIN_1_LOWER

// value minimum has 1 number (0-9)
utpass.MIN_1_NUMBER

// value minimum has 1 special character
utpass.MIN_1_SPECIAL

// default rules contains :
// - MIN 8 LENGTH
// - MIN 1 UPPER
// - MIN 1 LOWER
// - MIN 1 NUMBER
// - MIN 1 SPECIAL
utpass.DefaultRules()
```

### Available Languages

```go
// import utpass "github.com/kuromittsu/util_password"

// ID => Indonesian
utpass.LANG_ID

// EN => English
utpass.LANG_EN
```

### Structs

- Password Validate Result

```go
passwordValidateResult

// check validation valid or not
func(*passwordValidateResult) IsValid() bool

// get invalid list
// see Password Validate Invalid List
func(*passwordValidateResult) InvalidList() *passwordValidateInvalidList
```

- Password Validate Invalid List

```go
passwordValidateInvalidList

// get list of string error message
func(*passwordValidateInvalidList) Get() []string

// get error message but already join
func(*passwordValidateInvalidList) GetJoined(join string) (string, error)
```

### Usage

#### Example

```go
testList := []string{
  "#Password123",
  "#Password1",
  "#wwwpassword123",
  "Password",
  "21313123",
}

for _, v := range testList {
  result := utpass.PasswordValidate(v, utpass.DefaultRules(), utpass.LANG_EN)

  errMsgs, _ := result.InvalidList().GetJoined(", ")
  fmt.Printf("text: %s \n - is valid: %v \n - error message: %v \n", v, result.IsValid(), errMsgs)
}
```

Result

```txt
text: #Password123
 - is valid: true
 - error message:
text: #Password1
 - is valid: true
 - error message:
text: #wwwpassword123
 - is valid: false
 - error message: at least 1 uppercase letter
text: Password
 - is valid: false
 - error message: at least 1 number (0-9), at least 1 special character
text: 21313123
 - is valid: false
 - error message: at least 1 uppercase letter, at least 1 lowercase letter, at least 1 special character
```

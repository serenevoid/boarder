package common

/*
 Check error and panic if not nil.

 @param {error} err - error variable

 @example checkErr(err)
*/
func CheckErr(err error) {
    if err != nil {
        panic(err.Error())
    }
}

package ports

/** we should probably define our data schema in go
// I believe the protoc compiler will do that for us but not sure
// I will just use placeholder here
*/
 */

type Person struct {

}

type PersonId struct {

}

 type FindPersonKey struct {

}

type APIPort interface {
	AddNewPersonInfo(person Person) error
	EditPersonInfo(pid PersonId) error
	DeletePersonInfo(pid PersonId) error
	UpdatePersonInfo(pid PersonId) error
	FindPersonInfo(key FindPersonKey) (Person, error)
}

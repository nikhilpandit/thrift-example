// Thrift interface
// compile go code by running:
// thrift -r --gen go:thrift_import="github.com/apache/thrift/lib/go/thrift" hello.thrift

// Thrift structs define a common object. A struct has a set of strongly typed fields,
// each with a unique name identifier

struct Person {
	1: optional string id,
	2: optional string username,
	3: optional string first_name,
	4: optional string last_name, // simple string type
	5: optional i32 age, // integers can be i16, i32, i64
	6: optional list<string> hobbies, // example of list type
	7: optional Team team, // example of a field of custom type
}

struct Team {
	1: optional string name,
	2: optional bool active, // boolean type
}

exception HelloError {
	1: HelloErrorCode error_code,
	2: string error_message,
}

enum HelloErrorCode {
	NOT_FOUND = 1,
	OTHER_ERROR = 2,
}

// A service consists of a set of named functions, each with a list of parameters
// and a return type.
// Definition of a service is semantically equivalent to defining an interface in object oriented languages

/*
 * Hello service greets people
 */
service Hello {
	/*
	 * ping is used to make sure that the service is alive
	 */
	bool ping() throws (1: HelloError error),

	/*
	 * hello greets the person visiting the service
	 * @param username - a string username using which we can get the person
	 * @return string - greeting for the person
	 */
	string hello(1: string username) throws (1: HelloError error),
}


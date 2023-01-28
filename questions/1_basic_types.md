# Basic Types

## Type annotations and inference

A variable type can be defined by using **type annotations** or from **type inference**.

```ts
// explicitly define variable type usingtype annotations
const str: string = 'hello'; // define `str` as type `string`
const num: number = 123; // define `num` as type `number`
const bool: boolean = true; // define `bool` as type `boolean`

// variable type inferred from its declaration/initialization
const txt = 'world'; // `txt` defined as `string` (inferred from 'world')
const no = false; // `no` defined as `boolean` (inferred from false)
const nil = null; // `nil` defined as `null` (inferred from null)
```

## Primitives

- `string`: Any values that represent **string**.
- `number`: Any number. This includes `NaN`, **float**, **BigInt**.
- `boolean`: `true` and `false`.
- `null`: Value `null`.
- `undefined`: Value `undefined`.
- `symbol`: Represents a unique constant value.

```ts
// use type annotation to explicitly declare type
const str: string = 'foo';
const num: number = 123;

// when declaring variable, its type can be auto infered from its initiaized value
const isFresh = true; // inferred to type `boolean`
```

## Arrays

```ts
// both will create array of string
const arrStr1: string[] = ['a', 'b', 'c'];
const arrStr2: Array<string> = ['x', 'y', 'z'];

// array with mixed contents
const mixedArr: (string | boolean)[] = ['hello', 'foo', false];
const anyArr: any[] = ['baz', undefined, {}, false, null];
```

## Tupples

```ts
type Tuple2StringNumber = [string, number];
type Tuple3StringBoolNull = [string, boolean, null];

const tuple1: Tuple2StringNumber = ['abcd', 1234];
const tuple2: Tuple2StringNumber = ['abcd', 1234, 'efgh']; // Type '[string, number, string]' is not assignable to type 'Tuple2StringNumber'. Source has 3 element(s) but target allows only 2.

const tuple3: Tuple3StringBoolNull = ['qwerty', false, null];
const tuple4: Tuple3StringBoolNull = ['qwerty', false]; // Type '[string, false]' is not assignable to type 'Tuple3StringBoolNull'. Source has 2 element(s) but target requires 3.

// tuple can have optional
type TupleWithOptional = [string, string, string?];
const tuple5: TupleWithOptional = ['abcd', 'efgh'];
const tuple6: TupleWithOptional = ['abcd', 'efgh', 'ijkl'];
```

## Objects

```ts
const obj: {
  name: string;
  age: number;
  address: string;
  isHuman?: boolean; // optional property
} = {
  name: 'Luke',
  age: 99,
  address: 'Tatoine',
};
// object above can be initialized without property `isHuman` because it is optional
```

## Functions

A function parameters and return type can be defined. A function with return type `void` means it did not return any values. Please note that TypeScript's type `void` is different with JavaScript's `void` operator.

```ts
function functionName([param1: param1Type, [param2: param2Type, [...paramN: paramNType]]]): returnType {
  statement
}

function typedFunc(str: string, num: number, obj: { bool: boolean }): void {
  /*...*/
}

const arrowFunc = (str: string): boolean => {
  /*...*/

  // will error if return other than `boolean`
  // (including not returning anything)
  return true;
};

// return type can be more than one
function multipleReturn(): string | boolean {
  if (conditionA) {
    return true;
  }
  return 'foo';
}

// if the return type is not defined, the return type will be inferred
function implicitReturn(param: number | undefined) {
  return !!param; // this will implicitly set return type as `boolean`
}
```

### Function overloading

```ts
function add(a: string, b: string): string;

function add(a: number, b: number): number;

function add(a: any, b: any): any {
  return a + b;
}

add(4, 2);
add('4', '2');
add(4, '2');
// No overload matches this call.
//   Overload 1 of 2, '(a: string, b: string): string', gave the following error.
//     Argument of type 'number' is not assignable to parameter of type 'string'.
//   Overload 2 of 2, '(a: number, b: number): number', gave the following error.
//     Argument of type 'string' is not assignable to parameter of type 'number'.
```

## Interfaces

Interface is used to define an object-like type.

```ts
interface IAnimal {
  numOfLegs: number;
  isMammal: boolean;
}

// extending an interface
// adding properties to the already existing numOfLegs and isMammal
interface IAvian extends IAnimal {
  furColor: string;
  canFly: boolean;
  isMigrating?: boolean;
}

const cat: IAnimal = {
  numOfLegs: 4,
  isMammal: true,
};

const albatros: IAvian = {
  numOfLegs: 2,
  isMammal: false,
  furColor: 'white',
  canFly: true,
  isMigrating: true,
};

const quail: IAvian = {
  numOfLegs: 2,
  isMammal: false,
  furColor: 'brown',
  canFly: false,
};
```

## Type aliases

`type` can be used to create an alias for user defined type.

```ts
// literal type
type Foo = 'foo';
type One = 1;

let foo: Foo = 'foo';
foo = 'bar'; // Type '"bar"' is not assignable to type '"foo"'.
let one: One = 1;
one = 2; // Type '2' is not assignable to type '1'.

// alias for union types
type StringOrBool = string | boolean;
type NumOrNumArr = number | number[];
type NumFoo = number | 'foo';

let strBool: StringOrBool = 'foo';
strBool = null; // Type 'null' is not assignable to type 'StringOrBool'.

let numArr: NumOrNumArr = 123;
numArr = [1, 2, 3, 4];
numArr = 'bar'; // Type 'string' is not assignable to type 'NumOrNumArr'.
numArr = ['a', 3, 9]; // Type 'string' is not assignable to type 'number'.

let numFoo: NumFoo = 123;
numFoo = 456;
numFoo = 'foo';
numFoo = 'bar'; // Type '"bar"' is not assignable to type 'NumFoo'.

// function type
type NotEmptyString = (str: string) => boolean;
type PrintConsole = (inp: string | number) => void;

const notEmpty: NotEmptyString = (txt) => txt.trim().length > 0;
const printConsole: PrintConsole = (inp) => {
  let txt = inp;
  if (typeof inp === 'object') {
    txt = JSON.stringify(inp);
  }
  console.log(txt);
};

// alternative to interface
type Info = {
  detail: string;
  createdAt: Date;
  updatedAt: Date;
}

type User = {
  name: string;
  age: number;
  sex: 'male' | 'female';
  printInfo: PrintConsole;
  getInfo: (all?: boolean) => string;
  renewInfo(information: Info): void;
}
```

## Literal types

```ts
let changingString = "Hello World"; // Type string
// Because `changingString` can represent any possible string, that
// is how TypeScript describes it in the type system
 
const constantString = "Hello World"; // Type "Hello World"
// Because `constantString` can only represent 1 possible string, it
// has a literal type representation
```

```ts
type Review = 'worst' | 'bad' | 'good' | 'best';

const review1: Review = 'worst';
const review2: Review = 'bad';
const review3: Review = 'good';
const review4: Review = 'best';
const review5: Review = 'great'; // Type '"great"' is not assignable to type 'Review'.
```

```ts
function greeter(message: string, letterCase: 'upper' | 'lower') {
  switch(letterCase) {
    case 'upper':
      console.log(message.toUpperCase());
      break;
    default:
      console.log(message.toLowerCase());
      break;
  }
}

greeter("hEllO wOrLd!", 'lower');
```

## Enumerations

```ts
enum Agreement {
  No,
  Yes
}

enum Direction {
  Down,
  Up
}

const agreement = Agreement.Yes;
const direction = Direction.Up;
console.log(agreement, direction);
console.log(agreement === direction); // This condition will always return 'false' since the types 'Agreement' and 'Direction' have no overlap.
```

## Any

TypeScript's special type. Variable with this type can be assigned with any data type. It basically skip type checking.

```ts
let str: string;
let rand: any;

str = 123; // Type 'number' is not assignable to type 'string'.
str = false; // Type 'boolean' is not assignable to type 'string'.

rand = 'abc';
rand = 123;
rand = null;
```

## Unknown

```ts
let a: unknown = 'Hello world!';
let b: unknown = 42;
let c: unknown = false;

let aString: string = a as string;
let bNumber: number = b as number;
let cBoolean: boolean = c as boolean;

console.log(aString, bNumber, cBoolean);
```

```ts
interface IPerson {
  firstName: string;
  lastName: string;
}

let userInput: unknown = {
  firstName: 'Acong',
  lastName: 'Suherman'
};

function isPerson(input: unknown): input is IPerson {
  const inputPerson = input as IPerson;
  if (
    !('firstName' in inputPerson) ||
    !('lastName' in inputPerson)
  ) {
    return false;
  }
  if (
    typeof inputPerson.firstName !== 'string' ||
    typeof inputPerson.lastName !== 'string'
  ) {
    return false;
  }
  return true;
}

console.log(isPerson(userInput));
```

## Void

`void` type denotes the absence of having any type at all. It usually used as a return type for a function that doesnt return any value.

```ts
let a: void = undefined;
a = 1; // Type 'number' is not assignable to type 'void'.
a = null; // Type 'null' is not assignable to type 'void'.

function greeter(): void {
  console.log('Hello world!');
}
```

## Never

Type `never` is a type that "contains no value". A variable with type `never` cannot have any value. `never` usually used for a function that will always **throw an error**. The difference with type `void` when used on a function, function without return statement (or with return statement but no value) implicitly return `undefined`.

```ts
const str: never = 'abcd'; // Type 'string' is not assignable to type 'never'.

function returnNever1(): never {
  throw Error();
}

function returnNever2(): never {
  if (condition) {
    return 'ok'; // Type 'string' is not assignable to type 'never'.
  }
  throw new Error();
}
```

## References

- [Difference between 'object' ,{} and Object in TypeScript](https://stackoverflow.com/questions/49464634/difference-between-object-and-object-in-typescript)
- [Difference between string enums and string literal types in TS](https://stackoverflow.com/questions/49761972/difference-between-string-enums-and-string-literal-types-in-ts)
- ['unknown' vs. 'any'](https://stackoverflow.com/questions/51439843/unknown-vs-any)
- [What’s the difference between never and void in TypeScript?](https://medium.com/swlh/whats-the-difference-between-never-and-void-in-typescript-16f6629bfcdc)

## Exercise

[Character with longest consecutive repetition](https://www.codewars.com/kata/586d6cefbcc21eed7a001155/typescript)

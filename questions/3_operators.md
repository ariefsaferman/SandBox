# Operators and Modifiers

## AND `&` operator

The `&` operator is used to create an **intersection** type. In a more simple word, this operator create a type from a combination of two or more types. Combined type, has all properties of origin types. Please note that combining types that have similar property name but of different types will create an error.

```ts
interface ITaxonomy {
  family: string;
  genus: string;
  species: string;
}

interface IPhysiology {
  numLimbs: number;
  avgHeight: number;
  diet: 'omnivore' | 'carnivore' | 'herbivore';
}

type Organism = ITaxonomy & IPhysiology;

// type `Organism` has both `Taxonomy` and `Physiology` properties
const hooman: Organism = {
  avgHeight: 170,
  diet: 'omnivore',
  family: 'homonidae',
  genus: 'homo',
  numLimbs: 4,
  species: 'homo sapiens',
};

type User = {
  id: number;
  username: string;
  email: string;
};

type Shop = {
  id: string;
  shopname: string;
  shopAddress: string;
};

// `User` and `Shop` have property `id`, but different types
type UserShop = User & Shop;

const userShop: UserShop = {
  id: 1, // Type 'number' is not assignable to type 'never'.
  username: 'Acong',
  email: 'acong@mail.com',
  shopname: 'Sopi',
  shopAddress: 'Jl. Setiabudi'
}

// not possible
type StringNumber = string & number;
type NullUndef = null & undefined;
type StringArr = string & string[];
```

## OR `|` operator

The `|` operator is used to create a new type that combine other existing types. Unlike `&` operator that "merge" origin types, type constructed with `|` can have one of *any* origin types.

```ts
type StringNumber = string | number;
let strNum: StringNumber = 1;
strNum = 'one';
strNum = false; // Type 'boolean' is not assignable to type 'StringNumber'.

type User = {
  id: number;
  username: string;
  email: string;
};

type Shop = {
  id: string;
  shopname: string;
  shopAddress: string;
};

type UserShop = User | Shop;

const warung: UserShop = {
  id: 'shp001',
  shopname: 'Maju Jaya',
  shopAddress: 'Jl. Setiabudi',
  username: 'acong'
};

console.log(warung.username); // Property 'username' does not exist on type 'UserShop'. Property 'username' does not exist on type 'Shop'.

const pemilik: UserShop = {
  id: 123,
  username: 'acong',
  email: 'acong@mail.com',
  shopname: 'Sopi'
};

console.log(pemilik.shopname); // Property 'shopname' does not exist on type 'UserShop'. Property 'shopname' does not exist on type 'User'.
```

## Optional `?` operator

The `?` is to declare a property or a parameter as optional.

```ts
type Conspiracy = {
  desc: string;
  source?: string; // optional property
};

const lizardPeople: Conspiracy = {
  desc: 'Lizard people rules the world.'
};
```

```ts
function printConsole(txt: string, method?: 'log' | 'warn' | 'dir') {
  if (!method) {
    console.log(txt);
  } else {
    console[method](txt);
  }
}

printConsole('Hello world');
printConsole('Memory leak', 'warn');
```

## Non-null assertion `!` operator

```ts
interface IPerson {
  firstName: string;
  lastName?: string;
}

const person: IPerson = {
  firstName: 'Acong',
  lastName: 'Suherman'
};

console.log(person.lastName.toUpperCase()); // Object is possibly 'undefined'.

// I am sure that person.lastName is exist
console.log(person.lastName!.toUpperCase());

// without assertion
if (person.lastName) {
  console.log(person.lastName.toUpperCase());
}
```

## Optional chaining `?` operator

```ts
interface IPerson {
  firstName: string;
  lastName?: string;
}

const person: IPerson = {
  firstName: 'Djoko'
};

console.log(person.lastName?.toUpperCase()); // undefined
```

## Assertion `as` operator

An `as` can be used to **assert** a value to a more specific type. `as` can also be used to coerce a type into entirely different type.

```ts
interface IKeyValue {
  id: number;
  value: number;
}

let a = {
  id: 1,
  value: 10,
  description: 'lorem ipsum'
};

let b: IKeyValue = {
  id: 1,
  value: 10
};

a = b; // Property 'description' is missing in type 'KeyValue' but required in type '{ id: number; value: number; description: string; }'.

// I am sure that `a` has `KeyValue` type
let aa = a as IKeyValue;

aa = b;
```

## `readonly` modifier

The `readonly` modifier is used to make a property as a readonly.

```ts
interface IUser {
  readonly username: string;
  readonly email: string;
  password: string;
}

const acong: IUser = {
  name: 'acong',
  email: 'acong@mail.com',
  password: 'password',
};

acong.password = 'qwerty'; // ok
acong.name = 'djoko'; // error. cannot assign because read-only
```

## `typeof` operator

The `typeof` operator can be used to create a type based on other type (infer other type).

```ts
const sourceObj = {
  foo: 'foo',
  bar: 'baz',
  ok: false
};

type InferredObject = typeof sourceObj;

const imitationObj: InferredObject = {
  foo: 'oof',
  bar: 'zab',
  ok: true
};
```

```ts
const strNumArr = ['aaa', 33, 66, 'sss', 'rrr'];
type StringNumberArr = typeof strNumArr;
```

```ts
function add(num1: number, num2: number): number {
  return num1 + num2;
}

type BasicMath = typeof add;

const multiply: BasicMath = (num1, num2) => num1 * num2;
const subtract: BasicMath = (n1, n2) => n1 + n2;
```

## `keyof` operator

The `keyof` operator takes an object-like type and return its property names.

```ts
interface IFoo {
  lorem: string;
  ipsum: number;
  dolor: boolean;
}

type FooKeys = keyof IFoo;
let fooKey: FooKeys = 'dolor';
fooKey = 'ipsum';
fooKey = 'amet'; // Type '"amet"' is not assignable to type 'keyof Foo'.
```

## `in` operator

TypeScript's `in` operator behaves similar to JavaScript's `in`, that usually used as a safeguard when checking property name of an object. Usually used for **mapped types**.

```ts
// assume type `ILorem` as follows
interface ILorem {
  lorem: string;
  dolor: number;
  sit: Date;
  amet: () => void;
}

// this will create a "copy" of type `ILorem` with all of its properties
// as optional and type of string.
// see "mapped types" for more details/examples
type NeoLorem = {
  [L in keyof ILorem]?: string;
};

type Count = 'one' | 'two' | 'three' | 'four';

// type with properties 'one', 'two', 'three', and 'four'
type CountObject = {
  [C in Count]: number;
};
```

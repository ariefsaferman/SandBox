# Utility Types

TypeScript provides several utility types that usable for transformation. Below some types that we might use daily.

## `Partial<Type>`

Contructs a type with all properties as optional.

```ts
// all properties are required
interface IUser {
  name: string;
  dob: number;
  address: {
    country: string;
    zipCode: number;
  };
}

// make all `User`'s "first level" properties optional
type PartialUser = Partial<IUser>;

// properties `dob` and `address` are optional
const john: PartialUser = {
  name: 'John',
};

// although property `address` is optional, its property are not
// this will generate error since `address.country` is missing
const doe: PartialUser = {
  dob: 123456789,
  address: {
    zipCode: 7890,
  },
};

// since it is Partial, can initialized as empty object
const mrX: partialUser = {};
mrX.name = 'Mr. X';
```

## `Required<Type>`

The opposite of `Partial`, `Required` constructs a type where all properties are required.

```ts
interface IUser {
  name: string;
  dob?: number;
  sex?: 'male' | 'female';
  address?: {
    country?: string;
    zipCode?: number;
  };
}

type RequiredUser = Required<IUser>;

const jane: RequiredUser = {
  name: 'Jane Doe',
  dob: 1234567890,
  sex: 'female',
  // `address` is required, but its properties are still optional
  address: {
    country: 'NL',
  },
};
```

## `Readonly<Type>`

Construct a type where all properties are `readonly`. Selective read-only properties can be achieved by adding `readonly` **modifier** before property name.

```ts
// interface with partially read-only properties
interface IUser {
  readonly name: string;
  readonly dob: number;
  email: string;
  phoneNumber: number;
  isPremium: boolean;
}

type ReadOnlyUser = Readonly<IUser>;

const doe: IUser = {
  name: 'John Doe',
  dob: 223124512,
  email: 'doe_doe@mailer.com',
  phoneNumebr: 345678987,
  isPremium: false,
};

const wick: ReadOnlyUser = {
  name: 'John Wick',
  dob: 7793249235623,
  email: 'hollow_point@mailer.com',
  isPremium: false,
};

doe.name = 'John X. Doe'; // error. read-only
wick.name = 'John F. Wick'; // error. read-only

doe.email = 'the_real_john@mybusiness.com';
wick.email = 'at_your_service@theagency.com'; // error. read-only
```

## `Record<Key, Type>`

Create an object which `key` are type of `Key` and value are type of `Type`.

```ts
const a: Record<'lorem' | 'ipsum' | 'dolor', string> = {
  lorem: 'qwerty',
  // ipsum: 'aswd', // missing property `ipsum`
  dolor: true, // wrong (value) type
  foo: 'bar', // cannot assign `foo` 
};

type Job = 'mage' | 'knight' | 'thief';
const b: Record<Job, string | boolean> = {
  mage: 'Devastating magic',
  thief: 'Blazing fast movement',
  knight: true,
};
```

## `Pick<Type, Keys>`

Construct a type by picking one or more types from **interface**-like type.

```ts
interface IFox {
  lazy: string;
  fox: string;
  jump: string;
  over: string;
}

type PickedFox = Pick<IFox, 'fox' |  'jump'>;

const jumpingFox: PickedFox = {
  fox: 'Foxie',
  jump: 'wo...',
  lazy: 'its friday', // cannot assign `lazy`
  over: 'game over', // cannot assign `over`
};
```

## `Omit<Type, Keys>`

Contrary to `Pick`, `Omit` construct a type a removing several keys.

```ts
interface IFox {
  lazy: string;
  fox: string;
  jump: string;
  over: string;
}

type Omitted = Omit<IFox, 'fox' | 'jump'>;

const noFox: Omitted = {
  fox: 'Foxie', // cannot assign `fox`
  jump: 'wo...', // cannot assign `jump`
  lazy: 'its friday',
  over: 'game over',
};
```

## `Extract<Type, Union>`

Similar to `Pick` except for `Union` type.

```ts
const Numbers = 'one' | 'two' | 'three' | 'four';

type ExtractedNum = Extract<Numbers, 'one' | 'three'>;
let num1: Numbers = 'one';
let num2: ExtractedNum = 'one';
num1 = 'two';
num2 = 'two'; // error. cannot assign `two`
num1 = 'three';
num2 = 'three';
```

## `Exclude<Union, Excluded>`

Similar to `Omit`,  except for `Union` type.

```ts
const Numbers = 'one' | 'two' | 'three' | 'four';

type ExcludedNum = Exclude<Numbers, 'one' | 'three'>;
let num1: Numbers = 'one';
let num2: ExcludedNum = 'one';
num1 = 'two';
num2 = 'two';
num1 = 'three';
num2 = 'three'; // error. cannot assign `three`
```

## `Parameters<Type>`

To get parameter list (`arguments`) of a function.

```ts
type UserInfo = {
  name: string;
  dob: number;
}

type UserFunc = (name: string, usr: UserInfo) => UserInfo;

type FuncParams = Parameters<UserFunc>;
const funcAParam: FuncParams = ['Obi-wan Kenobi', { name: 'wick', dob: 12345 }];
const [usrName, usrInfo]: FuncParams = ['Obi-two Kenobi', { name: 'wick', dob: 67890 }];

const user: UserFunc = (name, info) => ({ ...info, name });

user(...funcAParam); // { name: 'Obi-wan Kenobi', dob: 12345 }
user(usrName, usrInfo); // { name: 'Obi-two Kenobi', dob: 67890 }

// since `arguments` is array-like, use parameter index as property name
const dob: FuncParams[1]['dob'] = 13579;
```

## `ReturnType<Type>`

Construct type of the return type of a **function** type.

```ts
function func1(): string | boolean {/*...*/}

// use `typeof` operator to get function type
let func1Return: ReturnType<typeof func1> = 'foo';
func1Return = false;
// func1Return = 123; // error

function func2(): { name: string; dob: number } {/*...*/}

const func2Return: ReturnType<typeof func2> = {
  name: 'john',
  dob: 876789,
};

// auto infer return type
function func3(m03: number, m05: boolean, def: string) {
  const c = Math.random();
  if (c > 0.3) {
    return m03;
  }
  if (c > 0.5) {
    return m05;
  }
  return def;
}

type Func3Return = ReturnType<typeof func3>; // string | number | boolean

// use defined function type
type Func4 = (txt: string) => { date: Date; days: number[] };

type Func4Return = ReturnType<Func4>;
const today: Func4Return['date'] = new Date();
const days: Func4Return['days'] = [3, 5, 2, 6];

// if the argument is not type of function
type ArgumentString = ReturnType<'abcd'>; // error
interface IDummy { date: Date; days: number[] };
type ArgumentArray = ReturnType<IDummy>; // error

```

## `Uppercase<Type>` and `Lowercase<Type>`

In a simple word, think it as ***String*** instance method `.toUpperCase()` and `.toLowerCase()`.

```ts
// with typed string
type UpperHello = 'HELLO WORLD';
type LowerIpsum = 'lorem ipsum';

const uh: UpperHello = 'HELLO WORLD';

const toUpper: Uppercase<LowerIpsum> = 'LOREM IPSUM';
const toLower: Lowercase<UpperHello> = 'hello world';
const randomHello: Lowercase<UpperHello> = 'Hello world'; // error

// with literal string
const litUpper: Uppercase<'Do not under estimate MY POWER'> = 'DO NOT UNDER ESTIMATE MY POWER';
const litLower: Lowercase<'Lazy Fox Jump Over Fence'>  = 'lazy fox jump over fence';
```

## `Capitalize<Type>` and `Uncapitalize<Type>`

Similar with `Uppercase` and `Lowercase`, except only for first character being checked (transformed). Next characters are left as is.

```ts
const capIpsum: Capitalize<'lOreM iPSum'> = 'LOreM iPSum';
const uncapIpsum: Uncapitalize<'LOreM IpsUm'> = 'lOreM IpsUm';
```
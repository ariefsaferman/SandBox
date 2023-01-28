# Generics

A type **generics** is a way to make defined type more usable, and defined (written) in a way similar to function.

For example, lets say we want to create an **object-like** type which has similar properties.

```ts
interface ICommonProps {
  length: number;
  isPublic: boolean;
  created_at: number;
  updated_at: number;
}

interface IStringData extends ICommonProps {
  get: (entry: string) => string;
  set: (...entries: string[]) => boolean;
}

interface INumberData extends ICommonProps {
  get: (entry: number) => number;
  set: (...entries: number[]) => boolean;
}
```

In example above, `IStringData` and `INumberData` have same properties with minor differences in data type. That type definition can be simplified by create a generic type (ex: `IData`) with one parameter `T`.

```ts
interface IData<T> extends ICommonProps {
  get: (entry: T) => T;
  set: (...entries: T[]) => boolean;
}

let stringData: Data<string>;
let numberData: Data<number>;
```

Generics can have default value too.

```ts
interface IData<T = string> extends ICommonProps {
  get: (entry: T) => T;
  set: (...entries: T[]) => boolean;
}

let stringData: Data;
let numberData: Data<number>;
```

More examples on generics.

```ts
function genericFunc<T>(inp: T) { }

genericFunc<number>('foo'); // Argument of type 'string' is not assignable to parameter of type 'number'.
genericFunc<boolean>(false);
genericFunc([1, true, 'foo']); // type infered from usage
```

```ts
const arrFunc = <A>(...input: A[]): A[] => {
  return input;
};
const arrFuncRet1 = arrFunc<boolean>(true, true, false); // boolean[]
const arrFuncRet2 = arrFunc('google.com', 'bing.com', 'duckduckgo.com'); // string[]
```

```ts
type Television<T> = (message: T) => T;
type Monitor<T> = (message: T, refreshRate: number) => T;
type Display<U, V> = Television<U> | Monitor<V>;
```

## Exercise

By using generics, implement [Bubble Sort Algoritm](https://www.geeksforgeeks.org/bubble-sort/) function which only accept array of strings or array of numbers as its input.

Hint: [Generic Constraints](https://www.typescriptlang.org/docs/handbook/2/generics.html#generic-constraints)
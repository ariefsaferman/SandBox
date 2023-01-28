# TypeScript Basic

## Introduction

### Why TypeScript?

- Each and every value in JavaScript has a set of behaviors you can observe from running different operations.
- JavaScript only truly provides dynamic typing - **running** the code to see what happens.
- Most people don't like to get any sorts of **errors when running** their code - those are considered bugs! And when we write new code, we try our best to avoid introducing new bugs.

```js
const message = 'Hello world!';

console.log(message.toLowerCase());
message(); // TypeError: message is not a function
```

```js
// Somewhere
function fn(x) {
  return x.flip();
}

// Executing fn() in your code
fn('Hello world!'); // TypeError: x.flip is not a function
```

### Initializing a TypeScript Project

```
npm init -y
npm i typescript
npx tsc --init
```

### Compiling (Transpiling) TypeScript Code to JavaScript Code

```
npx tsc
npx tsc --watch
```

## More

- [Basic Types](./1_basic_types.md)
- [Generics](./2_generics.md)
- [Operators](./3_operators.md)
- [Utility Types](./4_utility_types.md)

## References

- [TypeScript Deep Dive](https://basarat.gitbook.io/typescript/)
- [The TypeScript Handbook](https://www.typescriptlang.org/docs/handbook/intro.html)
- [What is a tsconfig.json](https://www.typescriptlang.org/docs/handbook/tsconfig-json.html)

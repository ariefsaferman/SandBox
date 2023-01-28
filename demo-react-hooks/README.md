# React Hooks with TypeScript

### Prerequisites:

- TypeScript

### References

- https://reactjs.org/docs/hooks-intro.html
- https://reactjs.org/docs/hooks-reference.html

### Table of Contents:

1.  React Hooks
2.  useState
3.  useEffect
4.  Custom Hooks



## 1. React Hooks

### 1.1. Definition

_Hooks_ are a new addition in React 16.8. They let you use state and other React features without writing a class.

**What is a Hook?** A Hook is a special function that lets you “hook into” React features. For example, `useState` is a Hook that lets you add React state to function components. We’ll learn other Hooks later.

### 1.1.2. Why Hooks
- Reusability
- Simplicity
- Clarity

Class Component with SetState
```javascript
class Toggle extends React.Component {
  constructor(props) {
    super(props);
    this.state = {isToggledOn: true};

    // This binding is necessary to make `this` work in the callback
    this.handleClick = this.handleClick.bind(this);
  }

  handleClick() {
    this.setState(prevState => ({
      isToggledOn: !prevState.isToggledOn
    }));
  }

  render() {
    return (
      <button onClick={this.handleClick}>
        {this.state.isToggledOn ? 'ON' : 'OFF'}
      </button>
    );
  }
}
```

Functional Component with Hooks
```javascript
function Toggle() {
  const [isToggledOn, setIsToggledOn] = useState(true);

  function handleClick() {
    setIsToggledOn(!isToggledOn)
  }

  return (
    <button onClick={handleClick}>
        {sToggledOn ? 'ON' : 'OFF'}
    </button>
  );
}
```

### 1.1.3. Hooks Rules
- Hooks can only be called inside React function components.
- Hooks can only be called at the top level of a component.
- Hooks cannot be conditional

### 1.2. Common Hooks:

- useState
- useEffect
- useContext
- useMemo
- useRef
- useReducer

Take a complete look at:
https://reactjs.org/docs/hooks-reference.html
(New) https://beta.reactjs.org/reference/react



## 2. useState

```typescript
const [data, setData] = useState<type>(initialData);
```

### 2.1. Definition

- `useState` allow us to track and modify state in a function component.
- Takes a parameter which is the **initial state value** or a function that return the state.
- Return 2 things, the **state** and a **function to update the state** (using what's called `destructuring` in JavaScript).
- During initial render (the first render), the returned state is the same as the `initialData`.
- The generic `<type>` can be omitted (removed) and React will detect the type based on `initialData` , but it is better to declare it for consistency.

### 2.2. Usage

1.  Import from react: `import { useState } from 'react'`
2.  Declare using format: `const [data, setData] = useState<type>(initialData)`
3.  Refer the state by calling `data`, and update the state by calling `setData(nextData)`
4.  We can update the state by passing a function that has a parameter of `previousState` and return the new state `setData((prevState) => prevState + 10)`
5.  If we want to start the initial data with `null` or `undefined` (like fetching data from server), we need to set the generic type with union:
    `const [data, setData] = useState<[] | null>(null)`

### 2.3. Example

#### 2.3.1. Simple String State

```typescript
// import first
import { useState } from "react"

// Inside a function component
const [name, setName] = useState<string>("")

// Refer inside JSX
<p>Hello {name}!</p>

// Change the data
setName("Budi")
```

#### 2.3.2. useState for Input tag on onChange handler

```typescript
// import first
import { useState } from "react"

// Inside a function component
const [name, setName] = useState<string>("")

// Refer inside JSX
<input value={name} onChange={handleChange} />

// create handleChange function to change input value
handleChange((e) => {
	setName(e.target.value)
})
```

#### 2.3.3. Array State

```typescript
// import first
import { useState } from "react"

// Inside a function component
const [arr, setArr] = useState<[] | null>(null)

// Refer inside JSX
<div>
	<ul>
		{arr.map((item, index) => (<li>Hello, you are number {item}!</li>))}
	</ul>
</div>

// Change the data
setArr[...arr, nextNumber]
```



## 3. useEffect

```typescript
useEffect(() => {
  // do something
  fetch(`${baseAPI}/products`).then((response) => response.JSON());
  //...
}, []);
```

### 3.1. Definition

- `useEffect` lets you perform side effects in function component.
- Example of side effects are data fetching, directly updating the DOM, and timers. Basically anything that takes time to execute.
- Takes 2 parameter; `EffectCallback` and `deps`
- `EffectCallbak` is a callback function that return a `void` or `Destructor` (a function that return `void` or `undefined`)
- `deps` is array of dependencies that the effect need to listen, to trigger executing the effect again
- If `deps` is empty array `[]`, effect will only run once

### 3.2. Usage

1. Import from react
   `import { useEffect } from 'react'`
2. Declare using format:
   `useEffect(() => { run something }, [arrayDependencies])`

### 3.3. Example

#### 3.3.1. Set Interval

```typescript
import { useState, useEffect } from "react";

function profile() {
  const [val, setVal] = useState(1);

  useEffect(() => {
    const timer = window.setInterval(() => {
      // setVal(val + 1);
      // because "closure" thing, `val` is always be '1'
      setVal((v) => v + 1);
    }, 1000);

    // need to clear the interval, but remember useEffect EffectCallback should only return void or Destructor (a function that return void | undefined)
    return () => window.clearInterval(timer);
  }, []);

  return <div>{val}</div>;
}
```

#### 3.3.2. Fetching Data from Server

```typescript
import { useState, useEffect } from "react";

interface IPost {
	userId: number;
	id: number;
	title: string;
	body: string;
}

function Posts() {
  const [post, setPost] = useState<IPost | undefined>(undefined);

  useEffect(() => {
		fetch("https://jsonplaceholder.typicode.com/posts/1")
			.then((response) => response.json())
			.then((result) => setPost(result))
			.catch((err) => console.log(err));
	}, []);

  return (
		<div>
			<h1>Posts</h1>
			<p>By: {post?.userId}</p>
			<p>By: {post?.body}</p>
		</div>
	);
}
```


## 4. Custom Hooks

### 4.1. Definition

- Custom Hooks lets you extract component logic into reusable functions
- Custom Hook is a JavaScript function whose name starts with `use` and that may call other Hooks
- Example Custom Hooks; `useFetch`, `useToggle`, `useAuth`

### 4.2. Usage

1. Create your own hooks file, for example `useFetch.tsx` (make sure to export it)
2. Import your custom hooks `import useFetch from '../hooks`
3. Call it as the way you initialized it, example `const { data, loading, error } = useFetch<Post[]>( "https://jsonplaceholder.typicode.com/posts" );`

### 4.3. Example

```typescript
// in useFetch.tsx
import { useEffect, useState } from "react";

export default function useFetch<Payload>(url: string): {
  data: Payload | null;
  loading: boolean | null;
  error: boolean;
} {
  const [data, setData] = useState<Payload | null>(null);
  const [loading, setLoading] = useState<boolean | null>(null);
  const [error, setError] = useState<boolean>(false);

  useEffect(() => {
    setLoading(true);
    fetch(url)
      .then((resp) => resp.json())
      .then((data: Payload) => setData(data))
      .catch((err) => {
        setError(true);
        console.log(err);
      })
      .finally(() => setLoading(false));
  }, [url]);

  return {
    data,
    loading,
    error,
  };
}

// in some component
import useFetch from "../hooks";

interface Post {
  userId: number;
  id: number;
  title: string;
  body: string;
}

export default function SomeComponent() {
  // Call the custom hooks, pass the generic type and assign from return values
  const { data, loading, error } = useFetch<Post[]>(
    "https://jsonplaceholder.typicode.com/posts"
  );

  return (
    <div>
      <p>Here is a custom hook 'useFetch'</p>
      {error && <em>Oh no, error :( please refresh the page</em>}
      {loading && <p>Loading...</p>}
      {!loading && (
        <ul>
          {data?.map((item, index) => (
            <li key={index}>
              {item.id}. {item.title}
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}
```

_Good luck! 😀_

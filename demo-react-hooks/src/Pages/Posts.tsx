import { useEffect, useState } from "react";

interface Post {
  userId: number;
  id: number;
  title: string;
  body: string;
}

export default function Posts() {
  const [posts, setPosts] = useState<Post[] | undefined>();
  // React Hooks Lifecycle Diagram
  // https://wavez.github.io/react-hooks-lifecycle/

  useEffect(() => {
    fetch("https://jsonplaceholder.typicode.com/posts")
      .then((response) => response.json())
      .then((result) => {
        setPosts(result);
      });

    console.log("this running after render");
  }, []);

  console.log("this running first");
  return (
    <div>
      <h1>Posts</h1>
      {posts?.map((post) => (
        <div key={post.id} className="comment">
          <p>By: {post.userId}</p>
          <p>{post.userId}</p>
          <p>{post.body}</p>
        </div>
      ))}
    </div>
  );
}

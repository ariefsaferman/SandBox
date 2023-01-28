import useFetch from "../hooks/useFetch";

interface Post {
  userId: number;
  id: number;
  title: string;
  body: string;
}

export default function PostsCustom() {
  const { data, loading, error } = useFetch<Post[]>(
    "https://jsonplaceholder.typicode.com/posts"
  );

  console.log("this running first");
  return (
    <div>
      <h1>Posts Custom</h1>

      {error && <em>Oh no, error :( please refresh the page</em>}
      {loading && <p>Loading...</p>}
      {!loading &&
        data?.map((post) => (
          <div key={post.id} className="comment">
            <p>By: {post.userId}</p>
            <p>{post.userId}</p>
            <p>{post.body}</p>
          </div>
        ))}
    </div>
  );
}

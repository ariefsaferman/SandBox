import useFetch from '../hooks/useFetch';

interface IPosts {
    userId: number;
    id: number;
    title: string;
    body: string;
}

const Posts = () => {
//    const [data] = useFetch<IPosts[]>('https://jsonplaceholder.typicode.com/posts');
   const {data, loading, error, errorMessage} = 
   useFetch<IPosts[]>('https://jsonplaceholder.typicode.com/posts');


    return (
        <div>
            <h1>Posts</h1>
            {loading && <p>Loading...</p>}
            {error && (
                <div>
                    <p>Oh no, error :) please refresh the page</p>
                    <p>Error message:</p>
                    <em>{errorMessage?.message}</em>
                </div>
            )}

            {data?.map((post) => {
                return <div key={post.id} className="comment">
                         <h2>By {post?.title}</h2>
                         <p>{post?.body}</p>
                        </div>    
            })}
        </div>
    )
}

export default Posts;
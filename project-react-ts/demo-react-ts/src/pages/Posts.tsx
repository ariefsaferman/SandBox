import React, {useState, useEffect} from 'react'; 

interface IPosts {
    userId: number;
    id: number;
    title: string;
    body: string;
}

const Posts = () => {
    // const [posts, setPosts] = useState<IPosts[]>([]);

    // useEffect(() => {
    //     fetch('https://jsonplaceholder.typicode.com/posts')
    //     .then(response => response.json())
    //     .then((result) => setPosts(result))
    //     .catch((err) => alert('Error fetching'))
    // }, [])

    
    return (
        <div>
            {/* <h1>Posts</h1>
            {posts?.map((post) => {
                return <div key={post.id} className="comment">
                         <h2>By {post?.title}</h2>
                         <p>{post?.body}</p>
                        </div>    
            })} */}
        </div>
    )
}

export default Posts;
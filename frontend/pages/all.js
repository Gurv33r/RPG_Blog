function Blog({ posts }) {
    return (
      <div className="posts-wrapper">
        {posts.map((post, i) => (
          <div className="post" key={i} >
            <div className="publish-date">{post.Date}</div>
            <p className="content">{post.Content}</p>
            <div className="update-time">{post.UpdatedAt}</div>
          </div>
        ))} 
      </div>
    )
  }

  // This function gets called at build time
export async function getStaticProps() {
    // Call an external API endpoint to get posts
    const res = await fetch('http://localhost:8080/all')
    const posts = await res.json();
    // By returning { props: { posts } }, the Blog component
    // will receive `posts` as a prop at build time
    return {
      props: {
         posts
      },
    }
  }
  
export default Blog
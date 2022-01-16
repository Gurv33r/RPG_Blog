function Blog({ posts }) {

  function jsxify(post, i) {
    let editDate = post.UpdatedAt
    if (!editDate.Valid){
      return (
        <div className="post" key={i} >
            <div className="publish-date">{post.Date}</div>
            <p className="content">{post.Content}</p>
        </div>
      )
    } else {
      return (
        <div className="post" key={i} >
          <div className="publish-date">{post.Date}</div>
          <p className="content">{post.Content}</p>
          <div className="edit-date">{editDate.Time}</div>
        </div>
      )
    }
  }

    return (
      <div className="posts-wrapper">
        {posts.map((post, i) => jsxify(post,i))} 
      </div>
    )
  }

export async function getStaticProps() {
    const res = await fetch('http://localhost:8080/all')
    const posts = await res.json();
    return {
      props: {
         posts
      },
    }
  }
  
export default Blog
import Link from "next/link"

function jsxify(post, i) {
  const editDate = post.UpdatedAt
  const publishDate = post.Date.split('T')[0]
  const jsx = [
    <div className="publish-date">{publishDate}</div>,
    <p className="content">{post.Content}</p>,
    <Link href={'/posts/'+publishDate}><u>See full post</u></Link>
  ]
  if (editDate){
    jsx.push(<div className="edit-date"> Edited on {editDate.Time.substring(0,10)} at {editDate.Time.substring(11,16)}</div>)
  }
  return  (
      <div className="post" key={i}>
        {jsx}
        <br/>
      </div>
  )
}

function Blog({ posts }) {
    return (
      <main>
        {posts.map((post, i) => jsxify(post,i))} 
      </main>
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
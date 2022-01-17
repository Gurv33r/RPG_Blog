import Link from 'next/link'

export default function Post( {post} ){
    const editDate = post.UpdatedAt
    const publishDate = post.Date.split('T')[0]
    const jsx = [
        <div className="publish-date">{publishDate}</div>,
        <p className="content">{post.Content}</p>,
        <Link href={`http://localhost:3000/posts/edit/${publishDate}`}><u>Edit this post</u></Link>
    ];
    if (editDate){
        jsx.push(<div className='edit-date'> Edited on {editDate.Time.substring(0,10)} at {editDate.Time.substring(11,16)}</div>)
    }
    return (
        <main>
            {jsx}
        </main>
    )   
}

export async function getStaticProps({params}){
    const res = await fetch(`http://localhost:8080/${params.date}`)
    if (res.status >= 500 ){
        alert(`Server error in retrieving post:${await res.text()}\nTry again later`)
    } else {
        const post = await res.json()
        return {
            props: {
                post
            },
        }
    }
}

export async function getStaticPaths(){
    const res = await fetch('http://localhost:8080/all')
    const posts = await res.json()
    const paths = posts.map((post) => ({
        params: {date: post.Date.split('T')[0]}
    }))
    return {
        paths,
        fallback: false
    }
}

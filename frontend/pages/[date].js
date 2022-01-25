import Head from 'next/head';
import Link from 'next/link'
import styles from '../styles/Home.module.css'
import postStyles from '../styles/Post.module.css'

export default function Post( {post} ){
    const editDate = post.UpdatedAt
    const publishDate = post.Date.split('T')[0]
    const jsx = [
        <div className={postStyles.publishDate}>{publishDate}</div>,
        <p className={postStyles.contentSpec}>{post.Content}</p>,
        <Link href={`/edit/${publishDate}`} onClick={(event)=>{
            event.preventDefault()
            document.body.style.cursor='progress'
        }}><u className={postStyles.linkSpec}>Edit this post</u></Link>
    ];
    if (editDate){
        jsx.push(<div className={postStyles.editDate}> Edited on {editDate.substring(0,10)} at {editDate.substring(11,16)}</div>)
    }
    return (
        <div className={styles.container}>
            <Head>
                <title>{publishDate}'s Post</title>
                <meta name='description' content={`Blog post published on ${publishDate}`}/>
            </Head>
            <main className={styles.main}>
                {jsx}
            </main>
        </div>
        
    )   
}

export async function getStaticProps({params}){
    const res = await fetch(`${process.env.BACKEND_URL}/${params.date}`)
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
    const res = await fetch(`${process.env.BACKEND_URL}/all`)
    const posts = await res.json()
    const paths = posts.map((post) => ({
        params: {date: post.Date.split('T')[0]}
    }))
    return {
        paths,
        fallback: false
    }
}

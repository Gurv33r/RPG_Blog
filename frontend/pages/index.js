import Link from "next/link"

import styles from '../styles/Home.module.css'
import postStyles from '../styles/Post.module.css';

import Head from "next/head"
import { useRouter } from "next/router";

function jsxify(post, i) {
  const editDate = post.UpdatedAt
  const publishDate = post.Date.split('T')[0]
  const jsx = [ 
  <div key={1} className={postStyles.publishDate}>
    {publishDate}
    </div>,
    <p key={2} className={postStyles.contentHome}>{post.Content.substring(0,10) + '...'}</p>,
    <Link key={3} href={`/${publishDate}`}><u className={postStyles.link}>See full post</u></Link>
  ]
  if (editDate){
    jsx.push(<div key={4} className={postStyles.editDate}>Edited on {editDate.substring(0,10)} at {editDate.substring(11,16)}</div>)
  }
  return  (
      <div className={postStyles.wrapper} key={i}>
        {jsx}
        <br/>
      </div>
  )
}

export default function Blog({ posts }) {
  const router = useRouter()
  if(posts){
    return (
      <div className={styles.container}>
        <Head>
          <title>RPG Blog</title>
          <meta name="description" content="Blog app made using React, Go, and PostgreSQL" />
        </Head>
        <main className={styles.main}>
          <button className={postStyles.btn} onClick={event => {
            event.preventDefault()
            router.push('/new')
          }}>New Post</button>
          {posts.map((post, i) => jsxify(post,i))} 
        </main>
        <footer className={styles.footer}>
          A Gurveer Singh Production &copy;
        </footer>
      </div>
    )
  }
  return (
    <div className={styles.container}>
        <Head>
          <title>RPG Blog</title>
          <meta name="description" content="Blog app made using React, Go, and PostgreSQL" />
        </Head>
        <main className={styles.main}>
          <button className={postStyles.btn} onClick={event => {
            event.preventDefault()
            router.push('/new')
          }}>New Post</button>
          <p>No posts so far!</p>
        </main>
        <footer className={styles.footer}>
          A Gurveer Singh Production &copy;
        </footer>
      </div>
  )
}

export async function getStaticProps() {
    const res = await fetch(`${process.env.BACKEND_URL}/all`)
    const posts = await res.json();
    return {
      props: {
         posts
      },
    }
  }
  
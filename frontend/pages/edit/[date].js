import { useRouter } from "next/router"

import styles from '../../styles/Form.module.css'

import Head from "next/head"

export default function Edit({ prevData }){
    const publishDate = prevData.Date.split('T')[0]
    const router = useRouter()
    async function submitForm(event){
        event.preventDefault()
        const res = await fetch(`${process.env.BACKEND_URL}/edit/${publishDate}`, {
            method:'POST',
            headers:{
                'Content-Type':'application/json'
            },
            body:JSON.stringify({
                'Date':prevData.Date,
                'Content':event.target.content.value,
            })
        })
        if(res.status === 404){
            router.push('/404')
        } else if (res.status === 400 || res.status === 500) {
            alert(`Error in sending message:${await res.text()}\nTry again later`)
            router.push('/')
        } else {
            alert(`Successfully edited ${publishDate}'s post`)
            router.push(`/${publishDate}`)
        }
    }

    return (
        <main>
            <Head>
                <title>{publishDate}&apos; Post</title>
                <meta name='description' content={`Blog post published on ${publishDate}`}/>
            </Head>
            <body>
                <div className={styles.form}>
                    <form onSubmit={submitForm}>
                        <h2>Date: {publishDate}</h2>
                        <textarea cols="1" rows="1" defaultValue={prevData.Content} name="content" className={styles.textarea}></textarea>
                        <button className={styles.btn} type="submit" id="submit-btn">Submit</button>
                        <button className={styles.btn} type="click" onClick={event => {
                            event.preventDefault()
                            router.push('/')
                        }}>Cancel</button>
                        <button className={styles.redBtn} type="click" onClick={deletePost}>Delete Post</button>
                    </form>  
                </div>
            </body>
        </main>
    )

    async function deletePost(event){
        event.preventDefault()
        const res = await fetch(`${process.env.BACKEND_URL}/remove/${publishDate}`)
        if (res.status === 500){
            alert(`Error in deleting post:${await res.text()}\nTry again later.`)
        } else {
            alert('Successfully deleted');
            router.push('/')
        }
    } 
}

export async function getStaticProps({params}) {
    const res = await fetch(`${process.env.BACKEND_URL}/${params.date}`)
    if (res.status >= 500){
        alert(`Server error in retrieving post:${await res.text()}\nTry again later`)
    } else {
        const prevData = await res.json()
        return {
            props: {
                prevData,
            },
        }
    }
}

export async function getStaticPaths(){
    const res = await fetch(`${process.env.BACKEND_URL}/all`)
    const posts = await res.json()
    const paths = posts.map(post => {
        params:{
            date: post.Date.split('T')[0]
        }
    })
    return {
        paths,
        fallback: false
    }
}
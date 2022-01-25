import { useRouter } from "next/router"
import styles from '../../styles/Form.module.css'
import Head from "next/head"

export default function NewPost( {dates} ){
    const router = useRouter()
    const currDate = new Date().toISOString()
    if (dates && dates.includes(currDate.split('T')[0])){
        const jsx = (
            <main>
                <Head>
                    <title>Problem with new post</title>
                    <meta name='description' content="Already posted today"/>
                </Head>
                <body>
                    <div className={styles.error}>
                        <h1>Houston, we have a problem!</h1>
                        <p className={styles.errMsg}>
                            It seems that you have already made a post for today&apos;s date!
                            Would you like to <b>edit</b> today&apos;s post?.
                        </p>
                        <button type="click" className={styles.btn} onClick={event => {
                            event.preventDefault()
                            router.push(`/edit/${currDate.split('T')[0]}`)
                        }} id="edit-post-btn">Yes</button>
                        <button className={styles.redBtn} type="click" onClick={event=>{
                            event.preventDefault()
                            router.push('/')
                        }}>No</button>
                    </div>    
                </body>
            </main>
        )
        return jsx
    }else {
        const submitForm = async event => {
            event.preventDefault()
            // assemble JSON object
            const data = {
                Date: currDate,
                Content: event.target.content.value
            }
            // send the JSON data 
            const res = await fetch(`${process.env.BACKEND_URL}/new`, {
                method: 'POST',
                headers:{
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            })
            if (res.status === 500){
                alert(`Error in submitting post:${await res.text()}\nTry again later!`)
                router.push('/')
            } else {
                alert('Successfully submitted')
                router.push(`/${data.Date.split('T')[0]}`)
            }
        }

        return (
            <main>
                <Head>
                    <title>New Post</title>
                    <meta name='description' content="New blog post form"/>
                </Head>
                <body>
                    <div className={styles.form}>
                        <h1>Today&apos;s date is {currDate.split('T')[0]}</h1>
                        <form onSubmit={submitForm}>
                            <textarea cols="1" rows="1" placeholder="Type out your thoughts here..." className={styles.textarea} name="content"></textarea>
                            <button className={styles.btn} type="submit">Submit</button>
                            <button className={styles.btn} type="click" onClick={event => {
                            event.preventDefault()
                            router.push('/')
                        }}>Cancel</button>
                        </form>
                    </div>
                </body>
            </main>
        )
    }
}

export async function getStaticProps(){
    const res = await fetch(process.env.BACKEND_URL + '/all')
    const data = await res.json()
    const dates = data ? data.map(post => {post.Date.split('T')[0]}) : null
    return{
        props:{
            dates: dates
        }
    } 
    
}
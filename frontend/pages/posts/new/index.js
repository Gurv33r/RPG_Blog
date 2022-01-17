import { useRouter } from "next/router"
import {dotenv} from 'dotenv'

export default function NewPost( {dates} ){
    const router = useRouter()
    const currDate = new Date().toISOString()
    if (!dates.includes(currDate.split('T')[0])){
        const jsx = (
            <main>
                <h1>Houston, we have a problem!</h1>
                <p className="error-msg">
                    It seems that you have already made a post for today's date!
                    If you want to <i>edit</i> today's post, click the button below.
                </p>
                <button type="click" onClick={router.push(`/edit/${currDate.split('T')[0]}`)} id="edit-post-btn">Edit today's post</button>
                <button type="click" onClick={router.go(-1)} id="go-back-button">No thanks, send me back</button>
            </main>
        )
    } else {
        const submitForm = async event => {
            event.preventDefault()
            // assemble JSON object
            const data = {
                Date: currDate,
                Content: event.target.content.value
            }
            // send the JSON data 
            const res = await fetch('http://localhost:8080/new', {
                method: 'POST',
                headers:{
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            })
            if (res.status === 500){
                alert(`Error in submitting post:${await res.text()}\nTry again later!`)
                router.push('/posts')
            } else {
                alert('Successfully submitted')
                router.push(`/posts/${data.Date}`)
            }
        }

        return (
            <main>
                <div className="form-wrapper">
                    <h1>Today's date is {currDate.split('T')[0]}</h1>
                    <form onSubmit={submitForm}>
                        <textarea cols="1" rows="1" placeholder="Type out your thoughts here..." classname="content-input" name="content"></textarea>
                        <button type="submit">Submit</button>
                    </form>
                </div>
            </main>
        )
    }
}

export async function getStaticProps(){
    dotenv.config()
    const res = await fetch(process.env.BACKEND_URL + '/all')
    const data = await res.json()
    const dates = data.map(post => {post.Date.split('T')[0]});
    return{
        props:{
            dates: dates
        }
    }
}
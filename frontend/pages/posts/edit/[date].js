import Link from "next/link";
import { useRouter } from "next/router";

export default function Edit({ prevData }){
    const publishDate = prevData.Date.split('T')[0]
    const router = useRouter()
    async function submitForm(event){
        event.preventDefault()
        const res = await fetch(`http://localhost:8080/edit/${publishDate}`, {
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
            router.push('/posts')
        } else {
            alert(`Successfully edited ${publishDate}'s post`)
            router.push(`/posts/${publishDate}`)
        }
    }

    return (
        <main>
        <h1>Add your edits below</h1>
            <form onSubmit={submitForm}>
                <h2>Date: {publishDate}</h2>
                <textarea cols="1" rows="1" id="new-content-input" defaultValue={prevData.Content} name="content"></textarea>
                <button type="submit" id="submit-btn">Submit</button>
                <Link href={'http://localhost:3000/posts'}>Cancel</Link>
                <button type="click" onClick={deletePost}>Delete Post</button>
            </form>  
        </main>
    )

    async function deletePost(event){
        event.preventDefault()
        const res = await fetch(`http://localhost:8080/remove/${publishDate}`)
        if (res.status === 500){
            alert(`Error in deleting post:${await res.text()}\nTry again later.`)
        } else {
            alert('Successfully deleted');
            router.push('/posts')
        }
    } 
}

export async function getStaticProps({params}) {
    const res = await fetch(`http://localhost:8080/${params.date}`)
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
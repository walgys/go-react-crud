import { useState } from 'react'
import './App.css'

function App() {
  const [name, setName] = useState('')
  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault()
    const options = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ name })
    };

    const response = await fetch('/users', options)
    const data = await response.json()

    console.log(data)
  }

  return (
    <>
      <div>
        <h1>Request data</h1>
        <button onClick={
          async () => {
            const response = await fetch('/users')
            const data = await response.json()
            console.log(data)
          }
        }>Get user data</button>
        </div>
        <div>
          <form action="/users" method="POST" onSubmit={handleSubmit}>
            <input type="text" name="username" onChange={event => setName(event.target.value)}/>
            <button type="submit">Create user</button>
          </form>
        </div>
    </>
  )
}

export default App

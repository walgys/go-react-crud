import './App.css'

function App() {

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
    </>
  )
}

export default App

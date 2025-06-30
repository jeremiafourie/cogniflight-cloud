import { useCallback, useState } from "react"

function App() {
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")

  const submitFn = useCallback(async (e) => {
    e.preventDefault()

    const res = await fetch("/api/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({email: email, pwd: password})
    });

    if(!res.ok) {
      if(res.status == 401) {
        console.log('Invalid credentials.')
      } else {
        console.log('Unknown response code: ', res.status)
      }
    }
  }, [email, password])

  return (
    <>
      <div style={{width: "100%", height: "100%", display: "flex", justifyContent: "center", alignItems: "center"}}>
        <div style={{border: "1px solid AccentColor", padding: "1rem 2rem"}}>
          <h4> Log in </h4>
          <hr/>
          <br/>

          <form style={{display: "grid", gridTemplateColumns: "1fr 1fr", gap: "1rem"}} onSubmit={submitFn}>
            <label htmlFor="email"> Email: </label>
            <input type="email" id="email" value={email} onChange={e => setEmail(e.target.value)} name="email"/>
            <label htmlFor="password"> Password: </label>
            <input type="password" id="password" value={password} onChange={e => setPassword(e.target.value)} name="password"/>

            <input type="submit" value="Log in"/>
          </form>
        </div>
      </div>
    </>
  )
}

export default App

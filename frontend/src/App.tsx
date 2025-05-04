import { Authenticator } from "@aws-amplify/ui-react";
import "./App.css";
import "@aws-amplify/ui-react/styles.css";

function App() {
  return (
    <Authenticator>
      {({ signOut, user }) => (
        <>
          <h1>Hello {user?.username}</h1>
          <button onClick={signOut}>Sign out</button>
        </>
      )}
    </Authenticator>
  );
}

export default App;

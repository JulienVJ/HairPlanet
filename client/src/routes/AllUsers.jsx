import { useState, useEffect } from 'react';
import "../styles/AllUsers.css"

function Users() {
  const [names, setNames] = useState([]);

  useEffect(() => {
    // Fetch data from the Golang API
    fetch('http://localhost:9192/users')
      .then(response => response.json())
      .then(data => {
        // Assuming the API returns an array of names
        setNames(data);
      })
      .catch(error => console.error('Error fetching data:', error));
  }, []); // Empty dependency array means this effect runs once after the first render

  return (
    <div>
      <h1 className='admin-title'>Liste des utilisateurs</h1>
      <div className='users-card'>
      {names.map((n) => (
        <>
        <div className='user-list'>
          <p key={n.key}>{n.name}</p>
          <p key={n.key}>{n.email}</p>
        </div>
        </>

      ))}        
      </div>
    </div>
  );
}

export default Users;

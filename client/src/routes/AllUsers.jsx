import { useState, useEffect } from 'react';
import "../styles/AllUsers.css"

function Users() {
  const [names, setNames] = useState([]);

  useEffect(() => {
    fetch('http://localhost:9192/users')
      .then(response => response.json())
      .then(data => {
        setNames(data);
      })
      .catch(error => console.error('Error fetching data:', error));
  }, []);

  return (
    <div>
      <h1 className='admin-title'>Liste des utilisateurs</h1>
      <div className='users-card'>
      {names.map((n) => (
        <>
        <div className='user-list'>
        <p key={n.key}>Nom : {n.lastName}</p>
          <p key={n.key}>Prénom : {n.firstName}</p>
          <p key={n.key}>Email : {n.email}</p>
        </div>
        </>

      ))}        
      </div>
    </div>
  );
}

export default Users;

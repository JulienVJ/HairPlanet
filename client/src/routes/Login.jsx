import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import "../styles/Home.css";

const Login = () => {
    const navigateTo = useNavigate();
    const [loginData, setLoginData] = useState({
        email: '',
        password: ''
    });

    const handleInputChange = (event) => {
        const { name, value } = event.target;
        setLoginData({
            ...loginData,
            [name]: value
        });
    };

    const handleSubmit = async (event) => {
        event.preventDefault();
        try {
            const response = await fetch('http://localhost:9192/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(loginData)
            });
            if (response.ok) {
                const data = await response.json();
                // Stockez le token d'authentification dans le stockage local ou les cookies
                localStorage.setItem('token', data.token);
                console.log('Login successful');
                navigateTo('/home'); // Redirigez l'utilisateur vers le tableau de bord après la connexion réussie
            } else {
                console.error('Login failed:', response.statusText);
                // Gérer l'échec de la connexion ici, par exemple, afficher un message d'erreur à l'utilisateur
            }
        } catch (error) {
            console.error('Error logging in:', error);
            // Gérer les erreurs réseau ici, par exemple, afficher un message d'erreur à l'utilisateur
        }
    };

    return (
        <div>
            <h2>Connexion</h2>
            <form onSubmit={handleSubmit}>
                <div>
                    <label htmlFor="email">Email:</label>
                    <input type="email" id="email" name="email" value={loginData.email} onChange={handleInputChange} required />
                </div>
                <div>
                    <label htmlFor="password">Mot de passe:</label>
                    <input type="password" id="password" name="password" value={loginData.password} onChange={handleInputChange} required />
                </div>
                <button type="submit">Se connecter</button>
            </form>
        </div>
    );
};

export default Login;

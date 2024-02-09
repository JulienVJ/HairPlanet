import { useState } from 'react';
import "../styles/Home.css"

const Login = () => {
    const [registrationData, setRegistrationData] = useState({
        email: '',
        password: '',
        isShop: false
    });

    const handleInputChange = (event) => {
        const { name, value } = event.target;
        setRegistrationData({
            ...registrationData,
            [name]: value
        });
    };

    const handleCheckboxChange = (event) => {
        const { name, checked } = event.target;
        setRegistrationData({
            ...registrationData,
            [name]: checked
        });
    };

    const handleSubmit = async (event) => {
        event.preventDefault();
        try {
            const response = await fetch('http://localhost:9192/register', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(registrationData)
            });
            if (response.ok) {
                console.log('Registration successful');
                // Gérer la réussite de l'inscription ici, par exemple, rediriger l'utilisateur vers une page de connexion
            } else {
                console.error('Registration failed:', response.statusText);
                // Gérer l'échec de l'inscription ici, par exemple, afficher un message d'erreur à l'utilisateur
            }
        } catch (error) {
            console.error('Error registering:', error);
            // Gérer les erreurs réseau ici, par exemple, afficher un message d'erreur à l'utilisateur
        }
    };

    return (
        <div>
            <h2>Inscription d'un utilisateur</h2>
            <form onSubmit={handleSubmit}>
                <div>
                    <label htmlFor="email">Email:</label>
                    <input type="email" id="email" name="email" value={registrationData.email} onChange={handleInputChange} required />
                </div>
                <div>
                    <label htmlFor="password">Mot de passe:</label>
                    <input type="password" id="password" name="password" value={registrationData.password} onChange={handleInputChange} required />
                </div>
                <div>
                    <label>
                        <input type="checkbox" name="isShop" checked={registrationData.isShop} onChange={handleCheckboxChange} />
                        Inscrire en tant que magasin
                    </label>
                </div>
                <button type="submit">S'inscrire</button>
            </form>
        </div>
    );
};

export default Login;
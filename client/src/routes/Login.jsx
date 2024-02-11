import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Button, TextField, Typography, Grid } from '@mui/material';


const Login = () => {
    const navigateTo = useNavigate();
    const [loginData, setLoginData] = useState({
        email: '',
        password: ''
    });
    
    const [errorMessage, setErrorMessage] = useState('');
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
                setErrorMessage("Votre email ou votre mot de passe semble être incorrect")
                // Gérer l'échec de la connexion ici, par exemple, afficher un message d'erreur à l'utilisateur
            }
        } catch (error) {
            console.error('Error logging in:', error);
            // Gérer les erreurs réseau ici, par exemple, afficher un message d'erreur à l'utilisateur
        }
    };
    

    return (
        <Grid container justifyContent="center" alignItems="center">
            <Grid item xs={12} sm={8} md={6} lg={4}>
                <Typography variant="h2">Connexion</Typography>
                <form onSubmit={handleSubmit}>
                    <TextField
                        id="email"
                        name="email"
                        label="Email"
                        type="email"
                        value={loginData.email}
                        onChange={handleInputChange}
                        fullWidth
                        required
                        margin="normal"
                    />
                    <TextField
                        id="password"
                        name="password"
                        label="Mot de passe"
                        type="password"
                        value={loginData.password}
                        onChange={handleInputChange}
                        fullWidth
                        required
                        margin="normal"
                    />
                     {errorMessage && <div style={{ color: 'red', padding: '20px' }}>{errorMessage}</div>}
                    <Button type="submit" variant="contained" color="primary">
                        Se connecter
                    </Button>
                </form>
            </Grid>
        </Grid>
    );
};

export default Login;

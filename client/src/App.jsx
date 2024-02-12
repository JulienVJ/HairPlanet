import './App.css';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Home from './routes/Home';
import Shop from './routes/Shop';
import FormHairdresser from './routes/FormHairdresser';
import Reservations from './routes/AllReservations';
import Users from './routes/AllUsers';
import Login from './routes/Login';
import Register from './routes/Register'

function App() {
  return (
    <Router>
      <Routes>
        <Route exact path="/" element={<Register />} />
        <Route exact path="/login" element={<Login />} />
        <Route exact path="/home" element={<Home />} />
        <Route path="/shop/:shopName" element={<Shop />} />
        <Route path="/createHairdresser" element={<FormHairdresser />} />
        <Route path="/users" element={<Users />} />
        <Route path="/reservations" element={<Reservations />} />
      </Routes>
    </Router>
  );
}

export default App;

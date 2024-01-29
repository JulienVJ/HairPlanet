import { useParams } from 'react-router-dom';

function Shop() {
    let { id } = useParams();

    return <h2>Shop ID: {id}</h2>;
}

export default Shop;

import { Link } from "react-router-dom";

const Home = () => {
    return (
        <div className="container mx-auto p-8">
            <h1 className="text-3xl font-bold text-center mb-4">Homepage</h1>
            <div className="flex justify-center space-x-4">
                <Link to={`order`} className="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded">Order</Link>
                <Link to={`bill`} className="bg-green-500 hover:bg-green-600 text-white font-bold py-2 px-4 rounded">Bill</Link>
            </div>
        </div>
    )
}

export default Home;
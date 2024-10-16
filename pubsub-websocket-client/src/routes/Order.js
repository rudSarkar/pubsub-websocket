import React, { useEffect, useState } from 'react';
import useWebSocket from 'react-use-websocket';

const Order = () => {
  const { sendMessage, lastMessage, readyState } = useWebSocket('ws://localhost:8080/ws');
  const [orders, setOrders] = useState([]);

  // Fetch initial order data on component mount
  useEffect(() => {
    const fetchOrders = async () => {
      try {
        const response = await fetch('http://localhost:8080/order'); // Update with your actual API endpoint
        const data = await response.json();
        setOrders(data);
      } catch (error) {
        console.error('Error fetching orders:', error);
      }
    };
    fetchOrders();
  }, []);

  // Send subscription message when connected
  useEffect(() => {
    if (readyState === 1) {
      sendMessage(JSON.stringify({ topic: 'orders' }));
    }
  }, [readyState, sendMessage]);

  // Handle incoming messages
  useEffect(() => {
    if (lastMessage) {
      const newOrder = JSON.parse(lastMessage.data);
      setOrders((prevOrders) => {
        const updatedOrders = [newOrder, ...prevOrders]; // Add new order at the beginning
        return updatedOrders.sort((a, b) => new Date(b.id) - new Date(a.id)); // Sort again
      });
    }
  }, [lastMessage]);

  return (
    <div className="p-4 max-w-3xl mx-auto bg-white shadow-md rounded-lg">
      <h1 className="text-2xl font-bold mb-4">Order Details</h1>
      {orders.length > 0 ? (
        <table className="min-w-full bg-white border border-gray-300">
          <thead>
            <tr>
              <th className="border-b-2 border-gray-300 px-4 py-2 text-left">Order ID</th>
              <th className="border-b-2 border-gray-300 px-4 py-2 text-left">Item</th>
              <th className="border-b-2 border-gray-300 px-4 py-2 text-left">Status</th>
            </tr>
          </thead>
          <tbody>
            {orders.map((order) => (
              <tr key={order.id}>
                <td className="border-b px-4 py-2">{order.id}</td>
                <td className="border-b px-4 py-2">{order.item}</td>
                <td className="border-b px-4 py-2">{order.Status}</td>
              </tr>
            ))}
          </tbody>
        </table>
      ) : (
        <p className="text-gray-500">No order data available.</p>
      )}
    </div>
  );
};

export default Order;
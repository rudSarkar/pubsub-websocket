import React from 'react';
import useWebSocket from 'react-use-websocket';

const Bill = () => {
  const { sendMessage, lastMessage, readyState } = useWebSocket('ws://localhost:8080/ws');

  // Send subscription message when connected
  React.useEffect(() => {
    if (readyState === 1) { // 1 means the WebSocket is open
      sendMessage(JSON.stringify({ topic: 'bills' }));
    }
  }, [readyState, sendMessage]);

  // Handle incoming messages
  const billData = React.useMemo(() => {
    if (lastMessage) {
      return JSON.parse(lastMessage.data);
    }
    return null;
  }, [lastMessage]);

  return (
    <div className="p-4 max-w-lg mx-auto bg-white shadow-md rounded-lg">
      <h1 className="text-2xl font-bold mb-4">Bill Details</h1>
      {billData ? (
        <div className="border border-gray-300 rounded-lg p-4">
          <div className="mb-2">
            <span className="font-semibold">Bill ID:</span> {billData.ID}
          </div>
          <div className="mb-2">
            <span className="font-semibold">Customer ID:</span> {billData.customer_id}
          </div>
          <div className="mb-2">
            <span className="font-semibold">Amount:</span> ${billData.amount.toFixed(2)}
          </div>
          <div className="mb-2">
            <span className="font-semibold">Created At:</span> {new Date(billData.CreatedAt).toLocaleString()}
          </div>
          <div className="mb-2">
            <span className="font-semibold">Updated At:</span> {new Date(billData.UpdatedAt).toLocaleString()}
          </div>
        </div>
      ) : (
        <p className="text-gray-500">No bill data available.</p>
      )}
    </div>
  );
};

export default Bill;

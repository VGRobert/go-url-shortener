import React, { useState } from 'react';

const UrlShortener = () => {
    const [url, setUrl] = useState('');
    const [shortUrl, setShortUrl] = useState('');
    const [error, setError] = useState(null);

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError(null); // Clear previous errors

        try {
            const response = await fetch('http://localhost:8080/shorten', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ url }),
            });

            if (!response.ok) {
                throw new Error('Failed to shorten the URL');
            }

            const data = await response.json();
            setShortUrl(`http://localhost:8080/${data.short_url}`);
        } catch (err) {
            setError(err.message);
        }
    };

    return (
        <div style={{ textAlign: 'center', marginTop: '50px' }}>
            <h1>URL Shortener</h1>
            <form onSubmit={handleSubmit}>
                <input
                    type="text"
                    value={url}
                    onChange={(e) => setUrl(e.target.value)}
                    placeholder="Enter URL to shorten"
                    style={{ width: '300px', padding: '10px', fontSize: '16px' }}
                />
                <button
                    type="submit"
                    style={{ padding: '10px 20px', marginLeft: '10px', fontSize: '16px' }}
                >
                    Shorten
                </button>
            </form>
            {error && (
                <div style={{ color: 'red', marginTop: '20px' }}>
                    <strong>Error:</strong> {error}
                </div>
            )}
            {shortUrl && (
                <div style={{ marginTop: '20px' }}>
                    <h2>Shortened URL:</h2>
                    <a href={shortUrl} target="_blank" rel="noopener noreferrer">
                        {shortUrl}
                    </a>
                </div>
            )}
        </div>
    );
};

export default UrlShortener;

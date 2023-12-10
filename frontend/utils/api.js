// THIS IS COPIED CODE; DO NOT USE BEYOND INSPIRATION

const BASE_URL = 'http://localhost:8080/api/';

// This function was modified by us to use the new API
export async function _fetchWithAuth(url, options = {}) {
    options.headers = options.headers || {
        'Content-Type': 'application/json',
        'credentials': 'include',
    };
    return fetch(url, {
        ...options,
        headers: {
            ...options.headers,
            uuid: `${localStorage.getItem('uuid')}`
        },
    });
}

// NOTHING BEYOND THIS LINE IS MODIFIED; DO NOT USE BEYOND INSPIRATION

function putAccessToken(token) {
    localStorage.setItem('accessToken', token);
}

function getAccessToken() {
    localStorage.getItem('accessToken');
}

export async function login({ formData }) {
    const response = await fetch(`${BASE_URL}/user/login`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            formData
        }),
    });

    const responseJson = await response.json();
    const { status, message } = responseJson;

    if (status !== 'success') {
        throw new Error(message);
    } else {
        console.log("Login successful");
    }

    const { data: { token } } = responseJson;
    return token;
}

async function getOwnProfile() {
    const response = await _fetchWithAuth(`${BASE_URL}/users/me`);

    const responseJson = await response.json();

    const { status, message } = responseJson;

    if (status !== 'success') {
        throw new Error(message);
    }

    const { data: { user } } = responseJson;

    return user;
}

async function getAllUsers() {
    const response = await fetch(`${BASE_URL}/users`);
    const responseJson = await response.json();

    const { status, message } = responseJson;

    if (status !== 'success') {
        throw new Error(message);
    }

    const { data: { users } } = responseJson;

    return users;
}

async function getAllTalks() {
    const response = await fetch(`${BASE_URL}/talks`);
    const responseJson = await response.json();

    const { status, message } = responseJson;

    if (status !== 'success') {
        throw new Error(message);
    }

    const { data: { talks } } = responseJson;

    return talks;
}

async function createTalk({ text, replyTo = '' }) {
    const response = await _fetchWithAuth(`${BASE_URL}/talks`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            text,
            replyTo,
        })
    });

    const responseJson = await response.json();
    const { status, message } = responseJson;

    if (status !== 'success') {
        throw new Error(message);
    }

    const { data: { talk } } = responseJson;
    return talk;
}

async function getTalkById(id) {
    const response = await fetch(`${BASE_URL}/talks/${id}`);
    const responseJson = await response.json();

    const { status, message } = responseJson;

    if (status !== 'success') {
        throw new Error(message);
    }

    const { data: { talkDetail } } = responseJson;

    return talkDetail;
}

async function toggleLikeTalk(id) {
    const response = await _fetchWithAuth(`${BASE_URL}/talks/likes`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            talkId: id,
        }),
    });

    const responseJson = await response.json();

    const { status, message } = responseJson;

    if (status !== 'success') {
        throw new Error(message);
    }
}

const api = {
    createTalk,
    getTalkById,
    putAccessToken,
    getAccessToken,
    getOwnProfile,
    register,
    login,
    toggleLikeTalk,
    getAllTalks,
    getAllUsers,
}

export default api;

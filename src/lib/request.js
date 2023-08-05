import axios from 'axios';

const postData = async({ endpoint, payload, onSuccess, onFailed }) => {
    // isi function.y apa
    const response = await axios({
        method: 'POST',
        url: `${import.meta.env.VITE_BASE_URL}${endpoint}`,
        data: payload
    });

    if (response.status === 200 || response.status === 201) {
        onSuccess(response)
    } else {
        onFailed(response)
    }
}

const postFormData = async({ endpoint, payload, onSuccess, onFailed }) => {
    const response = await axios({
        method: 'POST',
        url: `${import.meta.env.VITE_BASE_URL}${endpoint}`,
        data: payload,
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    });

    if (response.status === 200 || response.status === 201) {
        onSuccess(response.data.data.data)
    } else {
        onFailed(response)
    }
}

const putData = async({ endpoint, payload, onSuccess, onFailed }) => {
    // isi function.y apa
    const response = await axios({
        method: 'PUT',
        url: `${import.meta.env.VITE_BASE_URL}${endpoint}`,
        data: payload
    });

    if (response.status === 200) {
        onSuccess(response)
    } else {
        onFailed(response)
    }
}

const getData = async({ endpoint, params = {}, onSuccess, onFailed }) => {
    const response = await axios({
        method: 'GET',
        url: `${import.meta.env.VITE_BASE_URL}${endpoint}`,
        params: params
    })

    if (response.status === 200) {
        onSuccess(response.data.data.data)
    } else {
        onFailed(response)
    }
}

const deleteData = async({ endpoint, onSuccess, onFailed }) => {
    const response = await axios({
        method: 'DELETE',
        url: `${import.meta.env.VITE_BASE_URL}${endpoint}`,
    });

    if (response.status === 200 || response.status === 201) {
        onSuccess(response)
    } else {
        onFailed(response)
    }
}

export  { postData, postFormData, putData, deleteData, getData }
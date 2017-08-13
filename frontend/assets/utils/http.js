export function post (url, body) {
    const headers = {
        'X-Requested-With': 'XMLHttpRequest',
        'Content-Type':     'application/json',
        'Accept':           'application/json',
    };
    return fetch(
        url,
        {
            method: 'post',
            credentials: 'same-origin',
            redirect: 'manual',
            headers: headers,
            body: JSON.stringify(body)
        }
    );
}

export function del (url, body) {
    const headers = {
        'X-Requested-With': 'XMLHttpRequest',
        'Content-Type':     'application/json',
        'Accept':           'application/json',
    };
    return fetch(
        url,
        {
            method: 'delete',
            credentials: 'same-origin',
            redirect: 'manual',
            headers: headers,
            body: JSON.stringify(body)
        }
    );
}

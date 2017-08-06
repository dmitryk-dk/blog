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

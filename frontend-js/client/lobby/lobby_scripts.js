// Function to fetch the username.
async function fetchUsername() {
    try {
        const response = await fetch('/api/user/me');
        if (!response.ok) {
            throw new Error('Failed to fetch username');
        }
        const data = await response.json();
        const username = data.username;
        console.log('Welcome:', username);
        document.getElementById('username').textContent = `Get ready ${username}`;
        return username
    } catch (error) {
        console.error('Error fetching username:', error);
        return null;
    }
}

async function fetchLobbies() {
    try {
        const response = await fetch('/api/lobby/lobbies', {method: 'GET'},);
        if (!response.ok) {
            throw new Error('Failed to fetch lobbies');
        }
        return await response.json();
    } catch (error) {
        console.error('Error fetching lobbies:', error);
        return [];
    }
}

const createLobby = () => {
    //This Function creates the lobby when a name is inserted
    //and the user presses the create button
    const name = document.getElementById("lobby-text-box").value;
    const lobbyData = {
        lobby_name: name,
        uid: 1
    }

    fetch('/api/lobby/create', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(lobbyData)
    })
        .then(response => {
            if (!response.ok) {
                if (response.status === 405){
                    alert("Lobby limit reached, please delete a existing lobby")
                    return
                }
                alert("Failed to create lobby")
                return;
            }
            console.log('Lobby Created')
        })
        .catch(error => {
            console.error('Error creating lobby', error);
        }).finally(() => {
        location.reload()
    })
}

const deleteLobby = (lobbyID) => {
    fetch(`/api/lobby/remove/${lobbyID}`, {
        method: 'DELETE',
    })
        .then(response => {
            if (!response.ok) {
                console.log(response.statusText)
                alert('Failed to delete')
                return
            }
            console.log('Lobby deleted successfully');
            location.reload()
        })
        .catch(error => {
            console.error('Error deleting lobby:', error);
            alert('Failed to delete')
        }).finally(() => {
    });
}

function goToGame(lobbyUrl) {
    window.location = lobbyUrl
}

async function fetchLobbiesAndRender(currentUser) {
    try {
        //CONSTANTS INITIALIZED
        const lobbies = await fetchLobbies();
        const lobbyTable = document.createElement('table');
        const lobbyListContainer = document.querySelector('.lobby-container');
        const numRows = Math.ceil(lobbies.length / 4)
        lobbyListContainer.innerHTML = '';
        //NOW WE MAKE THE TABLE
        for (let i = 0; i < numRows; i++) {
            const row = lobbyTable.insertRow(); // Insert a new row
            for (let j = 0; j < 4; j++) {
                const index = i * 4 + j;
                if (index < lobbies.length) {
                    const cell = row.insertCell(); // Insert a new cell in the row
                    const lobby = lobbies[index];
                    const formedDate = new Date(lobby.createdAt).toLocaleString()

                    let deleteButtonHtml = ``
                    if (lobby.user === currentUser) {
                        deleteButtonHtml = `<button id="delete" class="join-button" onclick="deleteLobby(${lobby.lobbyId})">Delete</button>`;
                    }

                    cell.innerHTML = `
                          <div class="lobby-box">
                              <div class="lobby-name">
                                  <p class="small-text">${formedDate}</p>
                                  <hr>
                                  <p class="text">Lobby: ${lobby.lobbyName}</p>
                                  <p class="text">Owner: ${lobby.user}</p>
                              </div>
                              <div class="lobby-count">
                                  <p class="text">Players: ${lobby.joined}/4</p>
                              </div>
                              <div class="lobby-buttons">
                                ${deleteButtonHtml}                                
                                <button id="${lobby.lobbyId}-join" class="join-button" onclick="goToGame('/game?lobby=${lobby.lobbyId}')">Join</button>
                              </div>
                          </div>
                          `;
                }
            }
        }
        lobbyListContainer.appendChild(lobbyTable)
    } catch (error) {
        console.error('Error fetching and rendering lobbies:', error);
    }
}


function logout() {
    fetch('/logout', {
        method: 'GET',
    }).then(response => {
        if (response.redirected) {
            console.log('Redirecting to ' + response.url)
            window.location.href = response.url;
        } else {
            if (response.statusText) {
                alert('Error ' + response.statusText)
            }
        }
    }).catch(error => {
        console.error('There was a problem sending data to server:', error);
    });

}
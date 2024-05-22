// Initialize a Socket.IO connection
const socket = io.connect('http://localhost:8080');

socket.on('lobbyTimer', function(data) {
  document.getElementById(`${data.lobbyId}-timer`).innerHTML = ` <p class="text">Time Left to Join: ${data.timer}</p>`;
});

socket.on('inProgress', function(data) {
  document.getElementById(`${data.lobbyId}-timer`).innerHTML = '';
  document.getElementById(`${data.lobbyId}-progress`).innerHTML = ` <p class="text">Game in progress!</p>`;
  const joinButton = document.getElementById(`${data.lobbyId}-join`);
  
  // Remove the "join" button
  if (joinButton) {
    joinButton.parentNode.removeChild(joinButton);
  }
});

async function fetchLobbies() {
    try {
        const response = await fetch('/api/lobby');
        if (!response.ok) {
            throw new Error('Failed to fetch lobbies');
        }
        const lobbyResponse = await response.json();
        return lobbyResponse;

        // Extracting names and IDs from lobbies

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

    fetch('/api/lobby', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(lobbyData)
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('Failed to create lobby');
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
    fetch(`/api/lobby/${lobbyID}`, {
        method: 'DELETE',
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('Failed to delete lobby');
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
<template>
  <div class="group-list">
    <div class="group-list-header">
      <h2>{{ activeTab === "all" ? "All Groups" : "Your Groups" }}</h2>
      <div class="tab-navigation">
        <button
          @click="activeTab = 'all'"
          :class="{ active: activeTab === 'all' }"
        >
          All Groups
        </button>
        <button
          @click="activeTab = 'your'"
          :class="{ active: activeTab === 'your' }"
        >
          Your Groups
        </button>
      </div>
    </div>
    <div class="group-list-content">
      <ul>
        <li v-for="group in displayedGroups" :key="group.id" class="group-item">
          <div class="group-info">
            <h3>{{ group.name }}</h3>
            <p>{{ group.description }}</p>
          </div>
          <div class="group-actions">
            <button
              v-if="!hasJoined(group.id)"
              @click="joinGroup(group.id)"
              class="join-button"
            >
              Join
            </button>
            <div v-else>
              <!-- style the leave button red and if hover then dark red -->
              <button @click="leaveGroup(group.id)" class="leave-button">
                Leave
              </button>
              <button @click="enterGroup(group.id)" class="enter-button">
                Enter
              </button>
            </div>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      groups: [
        {
          id: 1,
          name: "Sergei Group",
          description: "Leave us alone, people are undesirable.",
        },
        { id: 2, name: "Sagar Group", description: "We like age of empires." },
        { id: 3, name: "Micke Group", description: "Sladda med bilen faaaan." },
        {
          id: 4,
          name: "Crippe Group",
          description: "Allt som händer i Vällingby.",
        },
        {
          id: 5,
          name: "Argan Group",
          description: "Join this group if you hate France and love anime.",
        },
        {
          id: 6,
          name: "Jay Group",
          description: "Free alcohol and sauna should be a human right.",
        },
        // Add more group objects here as needed
      ],
      activeTab: "all",
      userGroups: [1, 3, 5], // Example user's groups, replace with your own logic
    };
  },
  computed: {
    displayedGroups() {
      if (this.activeTab === "your") {
        return this.groups.filter((group) =>
          this.userGroups.includes(group.id)
        );
      }
      return this.groups;
    },
  },
  methods: {
    hasJoined(groupId) {
      return this.userGroups.includes(groupId);
    },
    joinGroup(groupId) {
      // Add logic to join the group
      this.userGroups.push(groupId);
    },
    leaveGroup(groupId) {
      // Add logic to leave the group
      const index = this.userGroups.indexOf(groupId);
      if (index > -1) {
        this.userGroups.splice(index, 1);
      }
    },
    enterGroup(groupId) {
      // Add logic to enter the group
      navigateTo("/groups/" + groupId);
      console.log("Entering group:", groupId);
    },
  },
};
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;600&display=swap");
.group-list {
  border: 1px solid #ccc;
  padding: 20px;
  border-radius: 4px;
  background-color: #f5f5f5;
  max-width: 500px;
  margin-bottom: 20px;
  height: 500px;
  font-family: "Poppins", sans-serif;
}

.group-list-header {
  margin-bottom: 16px;
}

.group-list-header h2 {
  font-size: 24px;
  font-weight: bold;
}

.tab-navigation {
  display: flex;
  margin-top: 10px;
}

.tab-navigation button {
  background: none;
  border: none;
  cursor: pointer;
  padding: 4px 8px;
  font-size: 14px;
  font-weight: bold;
  margin-right: 10px;
  color: #000000;
}

.tab-navigation button.active {
  border-bottom: 2px solid #000;
}

.group-list-content {
  max-height: 300px;
  overflow-y: auto;
  border: 1px solid #ccc;
}

.group-list-content ul {
  list-style-type: none;
  padding: 0;
  margin: 0;
}

.group-list-content li {
  margin-bottom: 16px;
  border-radius: 4px;
  padding: 12px;
  transition: transform 0.3s, box-shadow 0.3s;
}

.group-list-content li:hover {
  transform: translateY(-4px);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  background-color: white;
}

.group-list-content h3 {
  font-size: 18px;
  font-weight: bold;
  margin-bottom: 8px;
}

.group-list-content p {
  font-size: 14px;
  color: #888;
}

button {
  /* font-family: "Poppins", sans-serif; */
  color: #ffffff;
  margin: 10px;
}

.leave-button {
  background-color: rgb(255, 1, 1);
  color: white;
  border: none;
  color: #fff;
  font-size: 14px;
  cursor: pointer;
  padding: 8px 15px;
  border-radius: 4px;
}

.leave-button:hover {
  background-color: #de0101;
}

.join-button {
  background-color: #007bff;
  border: none;
  color: #fff;
  font-size: 14px;
  cursor: pointer;
  padding: 8px 15px;
  border-radius: 4px;
}

.join-button:hover {
  background-color: #0069d9;
}

.enter-button {
  background-color: #48ba0a;
  border: none;
  color: #fff;
  font-size: 14px;
  cursor: pointer;
  padding: 8px 15px;
  border-radius: 4px;
}

.enter-button:hover {
  background-color: #3d9708;
}

.group-item {
  display: flex;
  justify-content: space-between;
}

.group-info {
  flex: 1;
  margin-right: 10px;
}

.group-actions {
  display: flex;
  flex-direction: column;
}
</style>

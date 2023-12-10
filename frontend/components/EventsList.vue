<template>
  <div class="event-list">
    <div class="event-list-header">
      <h2>{{ activeTab === "all" ? "All Events" : "Your Events" }}</h2>
      <div class="tab-navigation">
        <button
          @click="activeTab = 'all'"
          :class="{ active: activeTab === 'all' }"
        >
          All Events
        </button>
        <button
          @click="activeTab = 'your'"
          :class="{ active: activeTab === 'your' }"
        >
          Your Events
        </button>
      </div>
    </div>
    <div class="event-list-content">
      <ul>
        <li v-for="event in displayedEvents" :key="event.id" class="event-item">
          <div class="event-info">
            <h3>{{ event.title }}</h3>
            <p>{{ event.description }}</p>
            <p>Day/Time: {{ event.dateTime }}</p>
          </div>
          <div class="event-actions">
            <button
              v-if="!hasRSVP(event.id)"
              @click="rsvpEvent(event.id)"
              class="accept-button"
            >
              Going
            </button>
            <div v-else>
              <button @click="cancelRSVP(event.id)" class="cancel-rsvp-button">
                Cancel
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
      events: [
        {
          id: 1,
          title: "Concert in the Park",
          description: "Enjoy live music in the park.",
          dateTime: "Saturday, July 15, 2023, 7:00 PM",
        },
        {
          id: 2,
          title: "Movie Night",
          description: "Watch a movie under the stars.",
          dateTime: "Friday, July 21, 2023, 8:30 PM",
        },
        {
          id: 3,
          title: "Art Exhibition",
          description: "Explore stunning artworks by local artists.",
          dateTime: "Sunday, July 30, 2023, 2:00 PM",
        },
      ],
      activeTab: "all",
      userRSVPs: [1, 3],
    };
  },
  computed: {
    displayedEvents() {
      if (this.activeTab === "your") {
        return this.events.filter((event) => this.userRSVPs.includes(event.id));
      }
      return this.events;
    },
  },
  methods: {
    hasRSVP(eventId) {
      return this.userRSVPs.includes(eventId);
    },
    rsvpEvent(eventId) {
      // Add logic to RSVP for the event
      this.userRSVPs.push(eventId);
    },
    cancelRSVP(eventId) {
      // Add logic to cancel the RSVP for the event
      const index = this.userRSVPs.indexOf(eventId);
      if (index > -1) {
        this.userRSVPs.splice(index, 1);
      }
    },
  },
};
</script>
<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;600&display=swap");
.event-list {
  border: 1px solid #ccc;
  padding: 20px;
  border-radius: 4px;
  background-color: #f5f5f5;
  max-width: 500px;
  margin-bottom: 20px;
  height: 500px;
  font-family: "Poppins", sans-serif;
}

.event-list-header {
  margin-bottom: 16px;
}

.event-list-header h2 {
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

.event-list-content {
  max-height: 300px;
  overflow-y: auto;
  border: 1px solid #ccc;
}

.event-list-content ul {
  list-style-type: none;
  padding: 0;
  margin: 0;
}

.event-list-content li {
  margin-bottom: 16px;
  border-radius: 4px;
  padding: 12px;
  transition: transform 0.3s, box-shadow 0.3s;
}

.event-list-content li:hover {
  transform: translateY(-4px);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  background-color: white;
}

.event-list-content h3 {
  font-size: 18px;
  font-weight: bold;
  margin-bottom: 8px;
}

.event-list-content p {
  font-size: 14px;
  color: #888;
}

button {
  color: #ffffff;
  margin: 10px;
}
.accept-button {
  background-color: #48ba0a;
  color: white;
  border: none;
  font-size: 14px;
  cursor: pointer;
  padding: 8px 15px;
  border-radius: 4px;
}

.accept-button:hover {
  background-color: #3f9e08;
}

.cancel-rsvp-button {
  background-color: rgb(255, 1, 1);
  color: white;
  border: none;
  font-size: 14px;
  cursor: pointer;
  padding: 8px 15px;
  border-radius: 4px;
}

.cancel-rsvp-button:hover {
  background-color: #de0101;
}

.event-item {
  display: flex;
  justify-content: space-between;
}

.event-info {
  flex: 1;
  margin-right: 10px;
}

.event-actions {
  display: flex;
  flex-direction: column;
}
</style>

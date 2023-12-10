export default defineNuxtPlugin(({ nuxtApp }) => {
  return {
    provide: {
      getEmailFromCookie: () => {
        const cookies = document.cookie;
        const cookieArray = cookies.split("; ");
        const cookieString = cookieArray.find((cookie) =>
          cookie.startsWith("NiceCookieBro=")
        );

        if (cookieString) {
          const encodedJsonString = cookieString.replace("NiceCookieBro=", "");
          const jsonString = decodeURIComponent(encodedJsonString);
          const cookieData = JSON.parse(jsonString);
          const email = cookieData.email;
          console.log("Email from cookie:", email);
        } else {
          console.error("Cookie not found");
        }
      },
      checkCookie: () => {
        const cookies = document.cookie;
        const cookieArray = cookies.split("; ");
        const cookieString = cookieArray.find((cookie) =>
          cookie.startsWith("NiceCookieBro=")
        );

        if (cookieString) {
          const encodedJsonString = cookieString.replace("NiceCookieBro=", "");
          const jsonString = decodeURIComponent(encodedJsonString);
          const cookieData = JSON.parse(jsonString);
          const email = cookieData.email;
          return email;
        } else {
          console.error("Cookie not found");
          navigateTo("/login");
        }
      },
    },
  };
});

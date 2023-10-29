Notification.requestPermission().then(permission => {
  if (permission === 'granted') {
    console.log('Notification permission granted.');
    const unread = 30;
    navigator.setAppBadge(unread);
  } else {
    console.error('Notification permission denied.');
  }
});
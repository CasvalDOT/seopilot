<script lang="ts">
	import { onDestroy, onMount } from "svelte";
	import Icon from "./icon.svelte";
	import { goto } from "$app/navigation";

  let timeRemaining: number = 0;
  let interval:number = -1;

  onMount(() => {
    interval = window.setInterval(() => {
      timeRemaining = (parseInt(window.localStorage.getItem("expiration_date") || "0") * 1000) - (new Date().getTime());
      if (timeRemaining <= 0) {
        clearInterval(interval);
        goto("/login");
      }
    },1000)
  })

  function numberEnding (value:number) {
      return (value > 1) ? 's' : '';
  }

  function millisecondsToStr (milliseconds:number) {
    var temp = Math.floor(milliseconds / 1000);

    var years = Math.floor(temp / 31536000);
    if (years) {
        return years + ' year' + numberEnding(years);
    }

    var days = Math.floor((temp %= 31536000) / 86400);
    if (days) {
        return days + ' day' + numberEnding(days);
    }
    var hours = Math.floor((temp %= 86400) / 3600);
    if (hours) {
        return hours + ' hour' + numberEnding(hours);
    }
    var minutes = Math.floor((temp %= 3600) / 60);
    if (minutes) {
        return minutes + ' minute' + numberEnding(minutes);
    }
    var seconds = temp % 60;
    if (seconds) {
        return seconds + ' second' + numberEnding(seconds);
    }
    return 'less than a second'; //'just now' //or other string you like;
}

onDestroy(() => {
  clearInterval(interval);
})

$: label = millisecondsToStr(timeRemaining);

</script>

<div 
  class:bg-error-500={timeRemaining < 5 * 60 * 1000}
  class:bg-warning-400={timeRemaining >= 5 * 60 * 1000 &&  timeRemaining < 15 * 60 * 1000}
  class:text-black={timeRemaining >= 5 * 60 * 1000 && timeRemaining < 15 * 60 * 1000}
  class:bg-surface-200-700-token={timeRemaining >= 12 * 60 * 1000}
  class="p-4 flex gap-x-4 items-center rounded-xl shadow-md"
>
  <Icon height="30" icon="solar:clock-circle-broken" />

  <div>
    <p><small>Session expires in</small></p>
    <p><strong>{label}</strong></p>
  </div>
</div>

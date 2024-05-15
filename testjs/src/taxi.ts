export const Taxi = (distance, minute) =>  {
    distance = Math.round(distance*2)/2
    minute = Math.ceil(minute)
    return (4*distance + minute)
}

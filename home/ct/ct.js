document.getElementById('contactForm').addEventListener('submit', function(event) {
    event.preventDefault(); // 폼의 기본 제출 동작을 막음

    // 입력 필드 값 가져오기
    const name = document.getElementById('name').value;
    const email = document.getElementById('email').value;
    const message = document.getElementById('message').value;

    // 폼 데이터 객체 생성
    const data = {
        name: name,
        email: email,
        message: message
    };

    // 서버로 POST 요청 보내기
    fetch('/submit', { // 여기에 실제 서버 URL 입력
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
    .then(response => response.json())
    .then(data => {
    console.log('Success:', data);
    // 성공 시 처리 로직 추가
    alert(`Form submitted successfully! ${data}님 소중한 의견 제출해주셔서 감사합니다`);
})

    .catch((error) => {
        console.error('Error:', error);
        // 오류 시 처리 로직 추가
        alert("An error occurred. Please try again.");
    });
});


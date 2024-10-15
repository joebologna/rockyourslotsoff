using System;
using System.Diagnostics;
using System.Drawing;
using System.IO;
using System.Linq;
using System.Windows.Forms;

public class MainForm : Form
{
    private string[] imagePaths =
    {
        "01-Apple.png",
        "02-Banana.png",
        "03-Blueberry.png",
        "04-Orange.png",
        "05-Strawberry.png",
        "06-Watermelon.png",
        "07-Seven.png",
    };

    private System.Windows.Forms.Timer timer;
    private System.Windows.Forms.Timer durationTimer;
    private Random random;
    private Button restartButton;
    private Label creditsLabel;
    private Label resultLabel;

    private int[] selectedIndices;
    private int credits;

    public MainForm()
    {
        this.Text = "Display PNG Images";
        this.Width = 800;
        this.Height = 600;
        this.FormBorderStyle = FormBorderStyle.FixedDialog; // Prevent resizing
        this.MaximizeBox = false; // Disable the maximize button

        random = new Random();

        timer = new System.Windows.Forms.Timer();
        timer.Interval = 100; // 100 ms
        timer.Tick += Timer_Tick;

        durationTimer = new System.Windows.Forms.Timer();
        durationTimer.Interval = 2000; // 2 seconds
        durationTimer.Tick += DurationTimer_Tick;

        selectedIndices = new int[imagePaths.Length]; // Initialize selectedIndices

        credits = 10; // Initialize credits

        creditsLabel = new Label
        {
            Text = $"Credits: {credits}",
            Location = new Point(10, 10),
            AutoSize = true
        };
        this.Controls.Add(creditsLabel);

        resultLabel = new Label
        {
            Text = "",
            Location = new Point(10, 40),
            AutoSize = true
        };
        this.Controls.Add(resultLabel);

        restartButton = new Button
        {
            Text = "Restart",
            Location = new Point(10, 70)
        };
        restartButton.Click += RestartButton_Click;
        this.Controls.Add(restartButton);
    }

    private int[] SpinReel(int length)
    {
        selectedIndices = new int[length];
        return selectedIndices.Select(_ => random.Next(0, length)).ToArray();
    }

    private void Timer_Tick(object? sender, EventArgs e)
    {
        var homePath = Environment.GetEnvironmentVariable("USERPROFILE");
        if (string.IsNullOrEmpty(homePath))
        {
            throw new InvalidOperationException("USERPROFILE environment variable is not set.");
        }

        Console.WriteLine("Reading files...");
        this.Controls.Clear(); // Clear previous images except the button and label
        this.Controls.Add(restartButton);
        this.Controls.Add(creditsLabel);
        this.Controls.Add(resultLabel);

        selectedIndices = SpinReel(3);

        for (int i = 0; i < selectedIndices.Length; i++)
        {
            var imagePath = imagePaths[selectedIndices[i]];
            var png = Path.Combine(homePath, "Projects\\2024\\rockyourslotsoff\\NETSlot\\Reel-Images", imagePath);
            try
            {
                PictureBox pictureBox = new PictureBox
                {
                    Image = Image.FromFile(png),
                    SizeMode = PictureBoxSizeMode.Zoom,
                    Width = 256,
                    Height = 256,
                    Margin = new Padding(0),
                    Location = new Point(i * 256 + 10, this.Height / 4 - 256 / 2)
                };
                this.Controls.Add(pictureBox);
            }
            catch (FileNotFoundException)
            {
                Debug.WriteLine($"File not found: {png}");
            }
            catch (Exception ex)
            {
                Debug.WriteLine($"Error loading image: {ex.Message}");
            }
        }
    }

    private void DurationTimer_Tick(object? sender, EventArgs e)
    {
        timer.Stop();
        durationTimer.Stop();
        var n = selectedIndices.Distinct().Count();
        Console.WriteLine($"n = {n}, selectedIndices = {string.Join(", ", selectedIndices)}");
        switch (n)
        {
            case 1:
                Console.WriteLine("You win big!");
                credits += 10;
                resultLabel.Text = "You win big!";
                break;
            case 2:
                Console.WriteLine("Push.");
                resultLabel.Text = "Push.";
                break;
            case 3:
                Console.WriteLine("You lose :(");
                credits -= 1;
                resultLabel.Text = "You lose :(";
                break;
        }
        creditsLabel.Text = $"Credits: {credits}";
    }

    private void RestartButton_Click(object? sender, EventArgs e)
    {
        StartTimers();
    }

    private void StartTimers()
    {
        timer.Start();
        durationTimer.Start();
    }

    [STAThread]
    public static void Main()
    {
        Application.EnableVisualStyles();
        Application.SetCompatibleTextRenderingDefault(false);
        Application.Run(new MainForm());
    }
}